package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/example/oauth2-go-single-entry/internal/config"
	"github.com/example/oauth2-go-single-entry/internal/httpapi"
	"github.com/example/oauth2-go-single-entry/internal/oauth"
	"github.com/example/oauth2-go-single-entry/internal/store"
	"github.com/example/oauth2-go-single-entry/internal/token"
)

func main() {
	cfg := config.Load()

	if err := os.MkdirAll("devkeys", 0700); err != nil {
		log.Fatal(err)
	}

	signingKey, err := token.LoadOrGenerateRSAKey(cfg.PrivateKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	users := store.NewUserStore()
	if err := users.SeedDemoUser("shailesh", "ChangeMe-123!"); err != nil {
		log.Fatal(err)
	}

	clients := store.NewClientStore()
	clients.Register(store.Client{
		ID:           cfg.ClientID,
		Secret:       cfg.ClientSecret,
		RedirectURI:  cfg.RedirectURI,
		AllowedScope: "profile orders:read",
	})

	codes := store.NewAuthorizationCodeStore()
	refreshTokens := store.NewRefreshTokenStore()
	revocations := store.NewRevocationStore()

	issuer := oauth.NewServer(oauth.ServerConfig{
		Issuer:          cfg.Issuer,
		ClientID:        cfg.ClientID,
		ClientSecret:    cfg.ClientSecret,
		RedirectURI:     cfg.RedirectURI,
		AccessTokenTTL:  cfg.AccessTokenTTL,
		RefreshTokenTTL: cfg.RefreshTokenTTL,
		PrivateKey:      signingKey,
		Users:           users,
		Clients:         clients,
		Codes:           codes,
		RefreshTokens:   refreshTokens,
		Revocations:     revocations,
	})

	api := httpapi.NewProtectedAPI(httpapi.ProtectedAPIConfig{
		Issuer:      cfg.Issuer,
		PrivateKey:  &signingKey.PublicKey,
		Revocations: revocations,
	})

	demoClient := httpapi.NewDemoClient(httpapi.DemoClientConfig{
		AuthBaseURL: cfg.Issuer,
		ClientID:    cfg.ClientID,
		Secret:      cfg.ClientSecret,
		RedirectURI: cfg.RedirectURI,
	})

	mux := http.NewServeMux()

	issuer.RegisterRoutes(mux)
	api.RegisterRoutes(mux)
	demoClient.RegisterRoutes(mux)

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		httpapi.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpapi.JSON(w, http.StatusOK, map[string]any{
			"service": "oauth2-go-single-entry",
			"status":  "running",
			"endpoints": []string{
				"/oauth/authorize",
				"/oauth/token",
				"/oauth/refresh",
				"/oauth/introspect",
				"/oauth/revoke",
				"/.well-known/oauth-authorization-server",
				"/.well-known/jwks.json",
				"/api/profile",
				"/api/orders",
				"/demo/login",
				"/healthz",
			},
		})
	})

	server := &http.Server{
		Addr:              cfg.Addr,
		Handler:           httpapi.SecurityHeaders(httpapi.RequestLogger(mux)),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("OAuth2 server listening on %s", cfg.Addr)
	log.Printf("Demo login: %s/demo/login", cfg.Issuer)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
