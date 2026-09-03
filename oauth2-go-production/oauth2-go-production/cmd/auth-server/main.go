package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/example/oauth2-go-production/internal/auth"
	"github.com/example/oauth2-go-production/internal/config"
	"github.com/example/oauth2-go-production/internal/httpserver"
	"github.com/example/oauth2-go-production/internal/store"
)

func main() {
	cfg := config.AuthConfigFromEnv()

	key, err := auth.LoadOrGenerateDevRSAKey(cfg.JWTPrivateKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	userStore := store.NewMemoryUserStore()
	codeStore := store.NewMemoryAuthorizationCodeStore()
	refreshStore := store.NewMemoryRefreshTokenStore()

	if err := userStore.SeedDemoUser("shailesh", "ChangeMe-123!"); err != nil {
		log.Fatal(err)
	}

	server := auth.NewServer(auth.ServerConfig{
		Issuer:          cfg.Issuer,
		ClientID:        cfg.ClientID,
		ClientSecret:    cfg.ClientSecret,
		RedirectURI:     cfg.RedirectURI,
		AccessTokenTTL:  cfg.AccessTokenTTL,
		RefreshTokenTTL: cfg.RefreshTokenTTL,
		PrivateKey:      key,
		UserStore:       userStore,
		CodeStore:       codeStore,
		RefreshStore:    refreshStore,
	})

	mux := http.NewServeMux()
	server.RegisterRoutes(mux)

	srv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           httpserver.SecurityHeaders(httpserver.RequestLogger(mux)),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("authorization server listening on %s", cfg.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	_ = os.Stdout
}
