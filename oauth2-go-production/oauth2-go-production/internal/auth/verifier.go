package auth

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/example/oauth2-go-production/internal/httpserver"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const claimsKey contextKey = "access_claims"

func ClaimsFromContext(ctx context.Context) AccessClaims {
	v, _ := ctx.Value(claimsKey).(AccessClaims)
	return v
}

type JWTVerifier struct {
	issuer string
	jwks   string
	mu     sync.RWMutex
	key    *rsa.PublicKey
}

func NewJWTVerifier(issuer, jwks string) (*JWTVerifier, error) {
	v := &JWTVerifier{issuer: issuer, jwks: jwks}
	if err := v.refresh(); err != nil {
		return nil, err
	}
	return v, nil
}

func (v *JWTVerifier) refresh() error {
	resp, err := http.Get(v.jwks)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var doc struct {
		Keys []struct {
			Kty string `json:"kty"`
			Kid string `json:"kid"`
			N   string `json:"n"`
			E   int    `json:"e"`
		} `json:"keys"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		return err
	}
	if len(doc.Keys) == 0 {
		return errors.New("no JWKS keys")
	}

	nBytes, err := base64.RawURLEncoding.DecodeString(doc.Keys[0].N)
	if err != nil {
		return err
	}

	pub := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: doc.Keys[0].E,
	}

	v.mu.Lock()
	v.key = pub
	v.mu.Unlock()
	return nil
}

func (v *JWTVerifier) RequireScope(scope string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			httpserver.JSON(w, http.StatusUnauthorized, map[string]string{
				"error": "missing_bearer_token",
			})
			return
		}

		raw := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))

		v.mu.RLock()
		key := v.key
		v.mu.RUnlock()

		var claims AccessClaims
		token, err := jwt.ParseWithClaims(raw, &claims, func(t *jwt.Token) (any, error) {
			if t.Method != jwt.SigningMethodRS256 {
				return nil, errors.New("unexpected signing method")
			}
			return key, nil
		},
			jwt.WithIssuer(v.issuer),
			jwt.WithAudience("resource-server"),
			jwt.WithExpirationRequired(),
		)

		if err != nil || !token.Valid {
			httpserver.JSON(w, http.StatusUnauthorized, map[string]string{
				"error": "invalid_token",
			})
			return
		}

		if claims.ExpiresAt == nil || time.Now().After(claims.ExpiresAt.Time) {
			httpserver.JSON(w, http.StatusUnauthorized, map[string]string{
				"error": "expired_token",
			})
			return
		}

		allowed := false
		for _, s := range strings.Fields(claims.Scope) {
			if s == scope {
				allowed = true
				break
			}
		}

		if !allowed {
			httpserver.JSON(w, http.StatusForbidden, map[string]string{
				"error": "insufficient_scope",
			})
			return
		}

		ctx := context.WithValue(r.Context(), claimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
