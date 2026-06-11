package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go-api-project/handlers"
)

func TestHealthHandler(t *testing.T) {

	req, err := http.NewRequest(
		"GET",
		"/health",
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(
		handlers.HealthHandler,
	)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf(
			"expected status 200 got %d",
			rr.Code,
		)
	}
}
