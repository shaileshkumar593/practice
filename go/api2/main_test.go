package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", getUsers)
	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("PUT /users", updateUser)

	return mux
}

func resetData() {
	mu.Lock()
	defer mu.Unlock()

	users = make(map[int]User)
	nextID = 1
}

func TestCreateUser(t *testing.T) {
	resetData()

	mux := setupMux()

	req := httptest.NewRequest(
		http.MethodPost,
		"/users",
		bytes.NewBuffer([]byte(`{"name":"john"}`)),
	)

	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected %d got %d",
			http.StatusCreated,
			rr.Code)
	}
}

func TestGetUsers(t *testing.T) {
	resetData()

	users[1] = User{
		ID:   1,
		Name: "john",
	}

	mux := setupMux()

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected %d got %d",
			http.StatusOK,
			rr.Code)
	}

	var result []User

	if err := json.NewDecoder(rr.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 user got %d",
			len(result))
	}
}

func TestUpdateUser(t *testing.T) {
	resetData()

	users[1] = User{
		ID:   1,
		Name: "john",
	}

	mux := setupMux()

	req := httptest.NewRequest(
		http.MethodPut,
		"/users?id=1",
		bytes.NewBuffer([]byte(`{"name":"rahul"}`)),
	)

	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected %d got %d",
			http.StatusOK,
			rr.Code)
	}

	if users[1].Name != "rahul" {
		t.Fatalf("expected rahul got %s",
			users[1].Name)
	}
}
