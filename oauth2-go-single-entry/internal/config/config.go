package config

import (
	"os"
	"time"
)

type Config struct {
	Addr            string
	Issuer          string
	PrivateKeyFile  string
	ClientID        string
	ClientSecret    string
	RedirectURI     string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func Load() Config {
	return Config{
		Addr:            env("APP_ADDR", ":8080"),
		Issuer:          env("OAUTH_ISSUER", "http://localhost:8080"),
		PrivateKeyFile:  env("JWT_PRIVATE_KEY_FILE", "./devkeys/private.pem"),
		ClientID:        env("OAUTH_CLIENT_ID", "demo-client"),
		ClientSecret:    env("OAUTH_CLIENT_SECRET", "dev-client-secret"),
		RedirectURI:     env("OAUTH_REDIRECT_URI", "http://localhost:8080/callback"),
		AccessTokenTTL:  duration("ACCESS_TOKEN_TTL", 10*time.Minute),
		RefreshTokenTTL: duration("REFRESH_TOKEN_TTL", 24*time.Hour),
	}
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func duration(key string, fallback time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if parsed, err := time.ParseDuration(value); err == nil {
			return parsed
		}
	}
	return fallback
}
