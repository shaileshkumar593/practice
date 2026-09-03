package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/example/oauth2-go-production/internal/httpserver"
	"github.com/example/oauth2-go-production/internal/store"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
	UserStore       *store.UserStore
	CodeStore       *store.AuthorizationCodeStore
	RefreshStore    *store.RefreshTokenStore
}

type Server struct {
	cfg ServerConfig
	mu  sync.RWMutex
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/oauth/authorize", s.authorize)
	mux.HandleFunc("/oauth/token", s.token)
	mux.HandleFunc("/.well-known/openid-configuration", s.discovery)
	mux.HandleFunc("/.well-known/jwks.json", s.jwks)
}

func (s *Server) authorize(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if q.Get("response_type") != "code" ||
		q.Get("client_id") != s.cfg.ClientID ||
		q.Get("code_challenge_method") != "S256" ||
		q.Get("code_challenge") == "" {
		http.Error(w, "invalid authorization request", http.StatusBadRequest)
		return
	}

	redirectURI := q.Get("redirect_uri")
	if redirectURI != s.cfg.RedirectURI {
		http.Error(w, "invalid redirect_uri", http.StatusBadRequest)
		return
	}

	if q.Get("scope") != "profile" {
		http.Error(w, "invalid scope", http.StatusBadRequest)
		return
	}

	state := q.Get("state")
	challenge := q.Get("code_challenge")

	if r.Method == http.MethodGet {
		httpserver.JSON(w, http.StatusOK, map[string]any{
			"message": "Demo authorization server",
			"login":   "POST credentials to /oauth/authorize",
			"state":   state,
			"note":    "For a real UI, render a login page and maintain an authenticated browser session.",
		})
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, ok := s.cfg.UserStore.Authenticate(username, password)
	if !ok {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	code, err := store.RandomToken(32)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	s.cfg.CodeStore.Put(store.AuthorizationCode{
		Code:          code,
		ClientID:      s.cfg.ClientID,
		UserID:        user.ID,
		RedirectURI:   redirectURI,
		Scope:         "profile",
		CodeChallenge: challenge,
		ExpiresAt:     time.Now().Add(60 * time.Second),
	})

	u, _ := url.Parse(redirectURI)
	params := u.Query()
	params.Set("code", code)
	params.Set("state", state)
	u.RawQuery = params.Encode()

	http.Redirect(w, r, u.String(), http.StatusFound)
}

func (s *Server) token(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	grantType := r.FormValue("grant_type")

	switch grantType {
	case "authorization_code":
		s.authorizationCodeToken(w, r)
	case "refresh_token":
		s.refreshTokenToken(w, r)
	default:
		tokenError(w, "unsupported_grant_type")
	}
}

func (s *Server) authorizationCodeToken(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	clientID := r.FormValue("client_id")
	redirectURI := r.FormValue("redirect_uri")
	verifier := r.FormValue("code_verifier")

	if clientID != s.cfg.ClientID || redirectURI != s.cfg.RedirectURI || verifier == "" {
		tokenError(w, "invalid_grant")
		return
	}

	c, ok := s.cfg.CodeStore.Consume(code)
	if !ok ||
		c.ClientID != clientID ||
		c.RedirectURI != redirectURI ||
		oauth2.S256ChallengeFromVerifier(verifier) != c.CodeChallenge {
		tokenError(w, "invalid_grant")
		return
	}

	access, err := s.issueAccessToken(c.UserID, c.Scope)
	if err != nil {
		tokenError(w, "server_error")
		return
	}

	refresh, err := store.RandomToken(48)
	if err != nil {
		tokenError(w, "server_error")
		return
	}

	s.cfg.RefreshStore.Put(store.RefreshToken{
		Token:     refresh,
		UserID:    c.UserID,
		ClientID:  c.ClientID,
		Scope:     c.Scope,
		ExpiresAt: time.Now().Add(s.cfg.RefreshTokenTTL),
	})

	httpserver.JSON(w, http.StatusOK, map[string]any{
		"access_token":  access,
		"token_type":    "Bearer",
		"expires_in":    int(s.cfg.AccessTokenTTL.Seconds()),
		"refresh_token": refresh,
		"scope":         c.Scope,
	})
}

func (s *Server) refreshTokenToken(w http.ResponseWriter, r *http.Request) {
	raw := r.FormValue("refresh_token")
	t, ok := s.cfg.RefreshStore.Rotate(raw)
	if !ok || t.ClientID != s.cfg.ClientID {
		tokenError(w, "invalid_grant")
		return
	}

	access, err := s.issueAccessToken(t.UserID, t.Scope)
	if err != nil {
		tokenError(w, "server_error")
		return
	}

	next, err := store.RandomToken(48)
	if err != nil {
		tokenError(w, "server_error")
		return
	}

	s.cfg.RefreshStore.Put(store.RefreshToken{
		Token:     next,
		UserID:    t.UserID,
		ClientID:  t.ClientID,
		Scope:     t.Scope,
		ExpiresAt: time.Now().Add(s.cfg.RefreshTokenTTL),
	})

	httpserver.JSON(w, http.StatusOK, map[string]any{
		"access_token":  access,
		"token_type":    "Bearer",
		"expires_in":    int(s.cfg.AccessTokenTTL.Seconds()),
		"refresh_token": next,
		"scope":         t.Scope,
	})
}

type AccessClaims struct {
	Scope string `json:"scope"`
	jwt.RegisteredClaims
}

func (s *Server) issueAccessToken(userID, scope string) (string, error) {
	now := time.Now()
	claims := AccessClaims{
		Scope: scope,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.cfg.Issuer,
			Subject:   userID,
			Audience:  jwt.ClaimStrings{"resource-server"},
			ExpiresAt: jwt.NewNumericDate(now.Add(s.cfg.AccessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        uuid.NewString(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = "dev-key-1"

	return token.SignedString(s.cfg.PrivateKey)
}

func (s *Server) discovery(w http.ResponseWriter, r *http.Request) {
	httpserver.JSON(w, http.StatusOK, map[string]any{
		"issuer":                           s.cfg.Issuer,
		"authorization_endpoint":           s.cfg.Issuer + "/oauth/authorize",
		"token_endpoint":                   s.cfg.Issuer + "/oauth/token",
		"jwks_uri":                         s.cfg.Issuer + "/.well-known/jwks.json",
		"response_types_supported":         []string{"code"},
		"grant_types_supported":            []string{"authorization_code", "refresh_token"},
		"code_challenge_methods_supported": []string{"S256"},
		"scopes_supported":                 []string{"profile"},
	})
}

func (s *Server) jwks(w http.ResponseWriter, r *http.Request) {
	pub := &s.cfg.PrivateKey.PublicKey
	jwk := map[string]any{
		"kty": "RSA",
		"use": "sig",
		"alg": "RS256",
		"kid": "dev-key-1",
		"n":   base64.RawURLEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(pub)),
		"e":   65537,
	}

	// For a real JWKS, n is the RSA modulus, not the DER public-key blob.
	nBytes := pub.N.Bytes()
	jwk["n"] = base64.RawURLEncoding.EncodeToString(nBytes)

	httpserver.JSON(w, http.StatusOK, map[string]any{
		"keys": []any{jwk},
	})
}

func tokenError(w http.ResponseWriter, code string) {
	httpserver.JSON(w, http.StatusBadRequest, map[string]string{
		"error": code,
	})
}

func LoadOrGenerateDevRSAKey(path string) (*rsa.PrivateKey, error) {
	if b, err := os.ReadFile(path); err == nil {
		block, _ := pem.Decode(b)
		if block == nil {
			return nil, errors.New("invalid PEM private key")
		}
		if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
			return key, nil
		}
		if keyAny, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
			if key, ok := keyAny.(*rsa.PrivateKey); ok {
				return key, nil
			}
		}
		return nil, errors.New("unsupported RSA private key")
	}

	if err := os.MkdirAll(filepathDir(path), 0700); err != nil {
		return nil, err
	}

	key, err := rsa.GenerateKey(rand.Reader, 3072)
	if err != nil {
		return nil, err
	}

	b := x509.MarshalPKCS1PrivateKey(key)
	pemData := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: b})

	if err := os.WriteFile(path, pemData, 0600); err != nil {
		return nil, err
	}
	return key, nil
}

func filepathDir(p string) string {
	i := strings.LastIndexAny(p, "/\\")
	if i < 0 {
		return "."
	}
	return p[:i]
}

func (s *Server) ClientSecret() string     { return s.cfg.ClientSecret }
func (s *Server) Context() context.Context { return context.Background() }

var _ = fmt.Sprintf
var _ = json.Valid
