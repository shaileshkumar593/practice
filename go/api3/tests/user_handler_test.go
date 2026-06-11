package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"api3/internal/handler"
	"api3/internal/service"
)

func setupRouter() *http.ServeMux {
	svc := service.NewUserService()

	h := handler.NewUserHandler(svc)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", h.GetUsers)
	mux.HandleFunc("POST /users", h.CreateUser)
	mux.HandleFunc("PUT /users", h.UpdateUser)

	return mux
}

func TestCreateUser(t *testing.T) {
	mux := setupRouter()

	body := []byte(`
	{
		"name":"john",
		"website":"https://google.com"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/users",
		bytes.NewBuffer(body),
	)

	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf(
			"expected %d got %d",
			http.StatusCreated,
			rr.Code,
		)
	}
}
