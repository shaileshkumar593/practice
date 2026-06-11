package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	service := NewUserService()
	handler := NewHandler(service)

	body := []byte(`{"name":"john"}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/users",
		bytes.NewBuffer(body),
	)

	w := httptest.NewRecorder()

	handler.CreateUser(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected %d got %d",
			http.StatusCreated,
			w.Code)
	}
}

func TestGetUsers(t *testing.T) {
	service := NewUserService()

	service.CreateUser("john")

	handler := NewHandler(service)

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	w := httptest.NewRecorder()

	handler.GetUsers(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected %d got %d",
			http.StatusOK,
			w.Code)
	}
}
