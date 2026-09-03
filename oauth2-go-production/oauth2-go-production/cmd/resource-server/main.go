package main

import (
	"log"
	"net/http"
	"time"

	"github.com/example/oauth2-go-production/internal/auth"
	"github.com/example/oauth2-go-production/internal/config"
	"github.com/example/oauth2-go-production/internal/httpserver"
)

func main() {
	cfg := config.ResourceConfigFromEnv()

	verifier, err := auth.NewJWTVerifier(cfg.Issuer, cfg.JWKSURL)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/profile", verifier.RequireScope("profile", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			claims := auth.ClaimsFromContext(r.Context())
			httpserver.JSON(w, http.StatusOK, map[string]any{
				"user_id": claims.Subject,
				"scope":   claims.Scope,
				"message": "protected resource accessed successfully",
			})
		},
	)))

	srv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           httpserver.SecurityHeaders(httpserver.RequestLogger(mux)),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("resource server listening on %s", cfg.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
