package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	appmiddleware "boilerplate/internal/middleware"

	kitlog "github.com/go-kit/log"
	"github.com/stretchr/testify/require"
)

func TestLogging(t *testing.T) {
	logger := kitlog.NewNopLogger()

	called := false

	next := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			called = true
			w.WriteHeader(http.StatusOK)
		},
	)

	handler := appmiddleware.Logging(next, logger)

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	require.True(t, called)
	require.Equal(t, http.StatusOK, rec.Code)
}
