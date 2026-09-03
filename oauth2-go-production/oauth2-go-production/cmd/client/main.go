package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"

	"github.com/example/oauth2-go-production/internal/config"
	"github.com/example/oauth2-go-production/internal/httpserver"
)

var (
	oauthConfig oauth2.Config
	verifier    string
	state       string
)

func main() {
	cfg := config.ClientConfigFromEnv()

	oauthConfig = oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       []string{"profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  cfg.AuthURL + "/oauth/authorize",
			TokenURL: cfg.AuthURL + "/oauth/token",
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/callback", callback)

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: httpserver.SecurityHeaders(httpserver.RequestLogger(mux)),
	}

	log.Printf("OAuth client listening on %s", cfg.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	_ = context.Background()
	_ = os.Stdout
}

func login(w http.ResponseWriter, r *http.Request) {
	var err error
	verifier = oauth2.GenerateVerifier()
	state, err = randomString(32)
	if err != nil {
		http.Error(w, "failed to create state", http.StatusInternalServerError)
		return
	}

	url := oauthConfig.AuthCodeURL(
		state,
		oauth2.S256ChallengeOption(verifier),
	)

	http.Redirect(w, r, url, http.StatusFound)
}

func callback(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("state") != state {
		http.Error(w, "invalid state", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "missing authorization code", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	token, err := oauthConfig.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		http.Error(w, "token exchange failed", http.StatusBadGateway)
		return
	}

	client := oauthConfig.Client(ctx, token)
	resp, err := client.Get("http://localhost:9100/api/profile")
	if err != nil {
		http.Error(w, "resource request failed", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "resource server rejected token", resp.StatusCode)
		return
	}

	httpserver.JSON(w, http.StatusOK, map[string]any{
		"access_token_received": true,
		"token_type":            token.TokenType,
		"expires_at":            token.Expiry,
		"resource_status":       resp.Status,
	})
}

func randomString(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

var _ = fmt.Sprintf
