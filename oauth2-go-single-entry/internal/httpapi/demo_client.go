package httpapi

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"sync"

	"golang.org/x/oauth2"
)

type DemoClientConfig struct {
	AuthBaseURL string
	ClientID    string
	Secret      string
	RedirectURI string
}

type DemoClient struct {
	cfg DemoClientConfig

	mu       sync.Mutex
	state    string
	verifier string
}

func NewDemoClient(cfg DemoClientConfig) *DemoClient {
	return &DemoClient{cfg: cfg}
}

func (c *DemoClient) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/demo/login", c.login)
	mux.HandleFunc("/callback", c.callback)
}

func (c *DemoClient) login(w http.ResponseWriter, r *http.Request) {
	verifier := oauth2.GenerateVerifier()

	state, err := randomString(32)
	if err != nil {
		http.Error(w, "failed to create state", http.StatusInternalServerError)
		return
	}

	c.mu.Lock()
	c.state = state
	c.verifier = verifier
	c.mu.Unlock()

	params := url.Values{}
	params.Set("response_type", "code")
	params.Set("client_id", c.cfg.ClientID)
	params.Set("redirect_uri", c.cfg.RedirectURI)
	params.Set("scope", "profile orders:read")
	params.Set("state", state)
	params.Set("code_challenge", oauth2.S256ChallengeFromVerifier(verifier))
	params.Set("code_challenge_method", "S256")

	http.Redirect(
		w,
		r,
		c.cfg.AuthBaseURL+"/oauth/authorize?"+params.Encode(),
		http.StatusFound,
	)
}

func (c *DemoClient) callback(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	expectedState := c.state
	verifier := c.verifier
	c.mu.Unlock()

	if r.URL.Query().Get("state") != expectedState {
		http.Error(w, "invalid state", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "missing authorization code", http.StatusBadRequest)
		return
	}

	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("client_id", c.cfg.ClientID)
	form.Set("client_secret", c.cfg.Secret)
	form.Set("redirect_uri", c.cfg.RedirectURI)
	form.Set("code_verifier", verifier)

	response, err := http.PostForm(
		c.cfg.AuthBaseURL+"/oauth/token",
		form,
	)
	if err != nil {
		http.Error(w, "token exchange failed", http.StatusBadGateway)
		return
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		http.Error(w, string(body), response.StatusCode)
		return
	}

	// This demo client makes one API call after exchanging the code.
	var tokenResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		TokenType    string `json:"token_type"`
	}

	if err := jsonUnmarshal(body, &tokenResponse); err != nil {
		http.Error(w, "invalid token response", http.StatusBadGateway)
		return
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		c.cfg.AuthBaseURL+"/api/profile",
		nil,
	)
	if err != nil {
		http.Error(w, "failed to build API request", http.StatusInternalServerError)
		return
	}

	req.Header.Set(
		"Authorization",
		tokenResponse.TokenType+" "+tokenResponse.AccessToken,
	)

	apiResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "resource request failed", http.StatusBadGateway)
		return
	}
	defer apiResponse.Body.Close()

	apiBody, _ := io.ReadAll(apiResponse.Body)

	HTML(w, http.StatusOK, fmt.Sprintf(`
<!doctype html>
<html>
<head><title>OAuth Success</title></head>
<body>
<h2>OAuth2 login successful</h2>
<p><b>Token type:</b> %s</p>
<p><b>Resource server status:</b> %s</p>
<h3>Protected API response</h3>
<pre>%s</pre>
<p><b>Refresh token was issued:</b> %t</p>
</body>
</html>`,
		HTMLEscape(tokenResponse.TokenType),
		HTMLEscape(apiResponse.Status),
		HTMLEscape(string(apiBody)),
		tokenResponse.RefreshToken != "",
	))
}

func randomString(size int) (string, error) {
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// Kept as a tiny wrapper to keep demo_client.go focused on HTTP flow.
func jsonUnmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func HTMLEscape(s string) string {
	return html.EscapeString(s)
}
