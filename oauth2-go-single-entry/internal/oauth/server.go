package oauth

import (
	"crypto/rsa"
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/example/oauth2-go-single-entry/internal/httpapi"
	"github.com/example/oauth2-go-single-entry/internal/store"
	"github.com/example/oauth2-go-single-entry/internal/token"
	"golang.org/x/oauth2"
)

type ServerConfig struct {
	Issuer          string
	ClientID        string
	ClientSecret    string
	RedirectURI     string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	PrivateKey      *rsa.PrivateKey
	Users           *store.UserStore
	Clients         *store.ClientStore
	Codes           *store.AuthorizationCodeStore
	RefreshTokens   *store.RefreshTokenStore
	Revocations     *store.RevocationStore
}

type Server struct {
	cfg ServerConfig
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/oauth/authorize", s.authorize)
	mux.HandleFunc("/oauth/login", s.login)
	mux.HandleFunc("/oauth/token", s.token)
	mux.HandleFunc("/oauth/refresh", s.refresh)
	mux.HandleFunc("/oauth/introspect", s.introspect)
	mux.HandleFunc("/oauth/revoke", s.revoke)
	mux.HandleFunc("/.well-known/oauth-authorization-server", s.discovery)
	mux.HandleFunc("/.well-known/jwks.json", s.jwks)
}

func (s *Server) authorize(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	client, ok := s.cfg.Clients.Get(q.Get("client_id"))
	if !ok || q.Get("response_type") != "code" {
		httpapi.OAuthError(w, http.StatusBadRequest, "invalid_request")
		return
	}

	if q.Get("redirect_uri") != client.RedirectURI {
		httpapi.OAuthError(w, http.StatusBadRequest, "invalid_request")
		return
	}

	if q.Get("code_challenge_method") != "S256" || q.Get("code_challenge") == "" {
		httpapi.OAuthError(w, http.StatusBadRequest, "invalid_request")
		return
	}

	if !scopeAllowed(q.Get("scope"), client.AllowedScope) {
		httpapi.OAuthError(w, http.StatusBadRequest, "invalid_scope")
		return
	}

	// Demo UI. A production authorization server should use a proper
	// authenticated browser session rather than accepting credentials
	// directly on this endpoint.
	httpapi.HTML(w, http.StatusOK, `
<!doctype html>
<html>
<head><title>OAuth Login</title></head>
<body>
<h2>Authorize application</h2>
<form method="post" action="/oauth/login">
<input type="hidden" name="client_id" value="`+htmlEscape(q.Get("client_id"))+`">
<input type="hidden" name="redirect_uri" value="`+htmlEscape(q.Get("redirect_uri"))+`">
<input type="hidden" name="state" value="`+htmlEscape(q.Get("state"))+`">
<input type="hidden" name="scope" value="`+htmlEscape(q.Get("scope"))+`">
<input type="hidden" name="code_challenge" value="`+htmlEscape(q.Get("code_challenge"))+`">
<label>Username <input name="username"></label><br><br>
<label>Password <input name="password" type="password"></label><br><br>
<button type="submit">Login and Authorize</button>
</form>
</body>
</html>`)
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpapi.OAuthError(w, http.StatusMethodNotAllowed, "invalid_request")
		return
	}

	client, ok := s.cfg.Clients.Get(r.FormValue("client_id"))
	if !ok || r.FormValue("redirect_uri") != client.RedirectURI {
		httpapi.OAuthError(w, http.StatusBadRequest, "invalid_request")
		return
	}

	user, ok := s.cfg.Users.Authenticate(
		r.FormValue("username"),
		r.FormValue("password"),
	)
	if !ok {
		httpapi.OAuthError(w, http.StatusUnauthorized, "access_denied")
		return
	}

	code, err := store.RandomToken(32)
	if err != nil {
		httpapi.OAuthError(w, http.StatusInternalServerError, "server_error")
		return
	}

	scope := r.FormValue("scope")

	s.cfg.Codes.Put(store.AuthorizationCode{
		Code:          code,
		ClientID:      client.ID,
		UserID:        user.ID,
		RedirectURI:   client.RedirectURI,
		Scope:         scope,
		CodeChallenge: r.FormValue("code_challenge"),
		ExpiresAt:     time.Now().Add(60 * time.Second),
	})

	redirect, err := url.Parse(client.RedirectURI)
	if err != nil {
		httpapi.OAuthError(w, http.StatusInternalServerError, "server_error")
		return
	}

	params := redirect.Query()
	params.Set("code", code)
	params.Set("state", r.FormValue("state"))
	redirect.RawQuery = params.Encode()

	http.Redirect(w, r, redirect.String(), http.StatusFound)
}

func (s *Server) token(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpapi.OAuthError(w, http.StatusMethodNotAllowed, "invalid_request")
		return
	}

	switch r.FormValue("grant_type") {
	case "authorization_code":
		s.authorizationCodeToken(w, r)
	case "refresh_token":
		s.refreshTokenToken(w, r)
	default:
		httpapi.OAuthError(w, http.StatusBadRequest, "unsupported_grant_type")
	}
}

func (s *Server) authorizationCodeToken(w http.ResponseWriter, r *http.Request) {
	client, ok := s.cfg.Clients.Get(r.FormValue("client_id"))
	if !ok || client.Secret != r.FormValue("client_secret") {
		httpapi.OAuthError(w, http.StatusUnauthorized, "invalid_client")
		return
	}

	if r.FormValue("redirect_uri") != client.RedirectURI {
		httpapi.OAuthError(w, http.StatusBadRequest, "invalid_grant")
		return
	}

	code, ok := s.cfg.Codes.Consume(r.FormValue("code"))
	if !ok {
		httpapi.OAuthError(w, http.StatusBadRequest, "invalid_grant")
		return
	}

	verifier := r.FormValue("code_verifier")
	if verifier == "" ||
		oauth2.S256ChallengeFromVerifier(verifier) != code.CodeChallenge {
		httpapi.OAuthError(w, http.StatusBadRequest, "invalid_grant")
		return
	}

	accessToken, expiresAt, err := s.issueAccessToken(code.UserID, code.Scope)
	if err != nil {
		httpapi.OAuthError(w, http.StatusInternalServerError, "server_error")
		return
	}

	refreshToken, err := store.RandomToken(48)
	if err != nil {
		httpapi.OAuthError(w, http.StatusInternalServerError, "server_error")
		return
	}

	s.cfg.RefreshTokens.Put(store.RefreshToken{
		Token:     refreshToken,
		UserID:    code.UserID,
		ClientID:  code.ClientID,
		Scope:     code.Scope,
		ExpiresAt: time.Now().Add(s.cfg.RefreshTokenTTL),
	})

	httpapi.JSON(w, http.StatusOK, map[string]any{
		"access_token":  accessToken,
		"token_type":    "Bearer",
		"expires_in":    int(time.Until(expiresAt).Seconds()),
		"refresh_token": refreshToken,
		"scope":         code.Scope,
	})
}

func (s *Server) refreshTokenToken(w http.ResponseWriter, r *http.Request) {
	client, ok := s.cfg.Clients.Get(r.FormValue("client_id"))
	if !ok || client.Secret != r.FormValue("client_secret") {
		httpapi.OAuthError(w, http.StatusUnauthorized, "invalid_client")
		return
	}

	old, ok := s.cfg.RefreshTokens.Rotate(r.FormValue("refresh_token"))
	if !ok || old.ClientID != client.ID {
		httpapi.OAuthError(w, http.StatusBadRequest, "invalid_grant")
		return
	}

	accessToken, expiresAt, err := s.issueAccessToken(old.UserID, old.Scope)
	if err != nil {
		httpapi.OAuthError(w, http.StatusInternalServerError, "server_error")
		return
	}

	next, err := store.RandomToken(48)
	if err != nil {
		httpapi.OAuthError(w, http.StatusInternalServerError, "server_error")
		return
	}

	s.cfg.RefreshTokens.Put(store.RefreshToken{
		Token:     next,
		UserID:    old.UserID,
		ClientID:  old.ClientID,
		Scope:     old.Scope,
		ExpiresAt: time.Now().Add(s.cfg.RefreshTokenTTL),
	})

	httpapi.JSON(w, http.StatusOK, map[string]any{
		"access_token":  accessToken,
		"token_type":    "Bearer",
		"expires_in":    int(time.Until(expiresAt).Seconds()),
		"refresh_token": next,
		"scope":         old.Scope,
	})
}

func (s *Server) refresh(w http.ResponseWriter, r *http.Request) {
	s.token(w, r)
}

func (s *Server) introspect(w http.ResponseWriter, r *http.Request) {
	raw := r.FormValue("token")
	if raw == "" || s.cfg.Revocations.IsRevoked(raw) {
		httpapi.JSON(w, http.StatusOK, map[string]any{"active": false})
		return
	}

	claims, err := token.ParseAndValidate(raw, &s.cfg.PrivateKey.PublicKey, s.cfg.Issuer)
	if err != nil {
		httpapi.JSON(w, http.StatusOK, map[string]any{"active": false})
		return
	}

	httpapi.JSON(w, http.StatusOK, map[string]any{
		"active":     true,
		"sub":        claims.Subject,
		"iss":        claims.Issuer,
		"aud":        claims.Audience,
		"scope":      claims.Scope,
		"exp":        claims.ExpiresAt.Unix(),
		"iat":        claims.IssuedAt.Unix(),
		"jti":        claims.ID,
		"token_type": "Bearer",
	})
}

func (s *Server) revoke(w http.ResponseWriter, r *http.Request) {
	raw := r.FormValue("token")
	if raw == "" {
		httpapi.OAuthError(w, http.StatusBadRequest, "invalid_request")
		return
	}

	// Demo behavior: revoke for the remaining access-token lifetime.
	// Production should identify token type and persist revocation safely.
	claims, err := token.ParseAndValidate(raw, &s.cfg.PrivateKey.PublicKey, s.cfg.Issuer)
	if err == nil && claims.ExpiresAt != nil {
		s.cfg.Revocations.Revoke(raw, claims.ExpiresAt.Time)
	} else {
		s.cfg.Revocations.Revoke(raw, time.Now().Add(s.cfg.RefreshTokenTTL))
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) issueAccessToken(userID, scope string) (string, time.Time, error) {
	return token.SignAccessToken(
		s.cfg.PrivateKey,
		s.cfg.Issuer,
		userID,
		scope,
		int64(s.cfg.AccessTokenTTL.Seconds()),
	)
}

func (s *Server) discovery(w http.ResponseWriter, r *http.Request) {
	httpapi.JSON(w, http.StatusOK, map[string]any{
		"issuer":                           s.cfg.Issuer,
		"authorization_endpoint":           s.cfg.Issuer + "/oauth/authorize",
		"token_endpoint":                   s.cfg.Issuer + "/oauth/token",
		"jwks_uri":                         s.cfg.Issuer + "/.well-known/jwks.json",
		"response_types_supported":         []string{"code"},
		"grant_types_supported":            []string{"authorization_code", "refresh_token"},
		"code_challenge_methods_supported": []string{"S256"},
		"scopes_supported":                 []string{"profile", "orders:read"},
	})
}

func (s *Server) jwks(w http.ResponseWriter, r *http.Request) {
	pub := &s.cfg.PrivateKey.PublicKey

	httpapi.JSON(w, http.StatusOK, map[string]any{
		"keys": []any{
			map[string]any{
				"kty": "RSA",
				"use": "sig",
				"alg": "RS256",
				"kid": token.KeyID,
				"n":   base64.RawURLEncoding.EncodeToString(pub.N.Bytes()),
				"e":   65537,
			},
		},
	})

}

func scopeAllowed(requested, allowed string) bool {
	allowedSet := map[string]bool{}
	for _, item := range strings.Fields(allowed) {
		allowedSet[item] = true
	}

	for _, item := range strings.Fields(requested) {
		if !allowedSet[item] {
			return false
		}
	}
	return true
}

func htmlEscape(value string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		`"`, "&quot;",
		"'", "&#39;",
	)
	return replacer.Replace(value)
}
