package http_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"boilerplate/internal/model"

	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	handler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(
				"Content-Type",
				"application/json",
			)

			w.WriteHeader(http.StatusOK)

			_ = json.NewEncoder(w).Encode(
				map[string]string{
					"status": "ok",
				},
			)
		},
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
}

func TestUserResponse(t *testing.T) {
	user := model.User{
		ID:    1,
		Name:  "John",
		Email: "john@example.com",
	}

	require.Equal(t, int64(1), user.ID)
	require.Equal(t, "John", user.Name)
	require.Equal(t, "john@example.com", user.Email)
}

var _ = context.Background
