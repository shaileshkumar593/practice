package httpapi

import (
	"context"
	"crypto/rsa"
	"errors"
	"net/http"
	"strings"

	"github.com/example/oauth2-go-single-entry/internal/store"
	"github.com/example/oauth2-go-single-entry/internal/token"
)

type ProtectedAPIConfig struct {
	Issuer      string
	PrivateKey  *rsa.PublicKey
	Revocations *store.RevocationStore
}

type ProtectedAPI struct {
	cfg ProtectedAPIConfig
}

func NewProtectedAPI(cfg ProtectedAPIConfig) *ProtectedAPI {
	return &ProtectedAPI{cfg: cfg}
}

type claimsKey struct{}

func (a *ProtectedAPI) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/api/profile", a.RequireScope("profile", http.HandlerFunc(a.profile)))
	mux.Handle("/api/orders", a.RequireScope("orders:read", http.HandlerFunc(a.orders)))
}

func (a *ProtectedAPI) RequireScope(scope string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := bearerToken(r)
		if raw == "" {
			JSON(w, http.StatusUnauthorized, map[string]string{
				"error": "missing_bearer_token",
			})
			return
		}

		if a.cfg.Revocations.IsRevoked(raw) {
			JSON(w, http.StatusUnauthorized, map[string]string{
				"error": "revoked_token",
			})
			return
		}

		claims, err := token.ParseAndValidate(
			raw,
			a.cfg.PrivateKey,
			a.cfg.Issuer,
		)
		if err != nil {
			JSON(w, http.StatusUnauthorized, map[string]string{
				"error": "invalid_token",
			})
			return
		}

		if !hasScope(claims.Scope, scope) {
			JSON(w, http.StatusForbidden, map[string]string{
				"error": "insufficient_scope",
			})
			return
		}

		ctx := context.WithValue(r.Context(), claimsKey{}, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *ProtectedAPI) profile(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(claimsKey{}).(token.Claims)
	if !ok {
		JSON(w, http.StatusInternalServerError, map[string]string{"error": "missing_claims"})
		return
	}

	JSON(w, http.StatusOK, map[string]any{
		"user_id": claims.Subject,
		"scope":   claims.Scope,
		"message": "profile endpoint accessed successfully",
	})
}

func (a *ProtectedAPI) orders(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(claimsKey{}).(token.Claims)
	if !ok {
		JSON(w, http.StatusInternalServerError, map[string]string{"error": "missing_claims"})
		return
	}

	JSON(w, http.StatusOK, map[string]any{
		"user_id": claims.Subject,
		"orders": []map[string]any{
			{"id": "ORD-1001", "status": "CONFIRMED"},
			{"id": "ORD-1002", "status": "PENDING"},
		},
	})
}

func bearerToken(r *http.Request) string {
	value := r.Header.Get("Authorization")
	if !strings.HasPrefix(value, "Bearer ") {
		return ""
	}
	return strings.TrimSpace(strings.TrimPrefix(value, "Bearer "))
}

func hasScope(scopeString, required string) bool {
	for _, scope := range strings.Fields(scopeString) {
		if scope == required {
			return true
		}
	}
	return false
}

var _ = errors.New
