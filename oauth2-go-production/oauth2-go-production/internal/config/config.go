package config

import (
	"os"
	"time"
)

type AuthConfig struct {
	Addr              string
	Issuer            string
	JWTPrivateKeyFile string
	ClientID          string
	ClientSecret      string
	RedirectURI       string
	AccessTokenTTL    time.Duration
	RefreshTokenTTL   time.Duration
}

func AuthConfigFromEnv() AuthConfig {
	return AuthConfig{
		Addr:              env("AUTH_ADDR", ":9000"),
		Issuer:            env("JWT_ISSUER", "http://localhost:9000"),
		JWTPrivateKeyFile: env("JWT_PRIVATE_KEY_FILE", "./devkeys/private.pem"),
		ClientID:          env("CLIENT_ID", "demo-web"),
		ClientSecret:      env("CLIENT_SECRET", "dev-client-secret-change-me"),
		RedirectURI:       env("CLIENT_REDIRECT_URL", "http://localhost:8080/callback"),
		AccessTokenTTL:    duration("ACCESS_TOKEN_TTL", 10*time.Minute),
		RefreshTokenTTL:   duration("REFRESH_TOKEN_TTL", 24*time.Hour),
	}
}

type ResourceConfig struct {
	Addr    string
	Issuer  string
	JWKSURL string
}

func ResourceConfigFromEnv() ResourceConfig {
	return ResourceConfig{
		Addr:    env("RESOURCE_ADDR", ":9100"),
		Issuer:  env("JWT_ISSUER", "http://localhost:9000"),
		JWKSURL: env("JWKS_URL", "http://localhost:9000/.well-known/jwks.json"),
	}
}

type ClientConfig struct {
	Addr         string
	AuthURL      string
	ResourceURL  string
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func ClientConfigFromEnv() ClientConfig {
	return ClientConfig{
		Addr:         env("CLIENT_ADDR", ":8080"),
		AuthURL:      env("AUTH_URL", "http://localhost:9000"),
		ResourceURL:  env("RESOURCE_URL", "http://localhost:9100"),
		ClientID:     env("CLIENT_ID", "demo-web"),
		ClientSecret: env("CLIENT_SECRET", "dev-client-secret-change-me"),
		RedirectURL:  env("CLIENT_REDIRECT_URL", "http://localhost:8080/callback"),
	}
}

func env(k, fallback string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return fallback
}

func duration(k string, fallback time.Duration) time.Duration {
	if v := os.Getenv(k); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return fallback
}
