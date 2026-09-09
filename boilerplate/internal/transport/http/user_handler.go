package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"boilerplate/internal/endpoint"

	userreq "boilerplate/internal/request/user"
	userresp "boilerplate/internal/response/user"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewUserHandler(e endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)

	getUser := kithttp.NewServer(
		e.GetUser,
		decodeGetUserRequest,
		encodeGetUserResponse,
	)

	createUser := kithttp.NewServer(
		e.CreateUser,
		decodeCreateUserRequest,
		encodeCreateUserResponse,
	)

	r.Handle("/v1/users/{id}", getUser).Methods(http.MethodGet)
	r.Handle("/v1/users", createUser).Methods(http.MethodPost)

	return r
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil || id <= 0 {
		return nil, errors.New("invalid user id")
	}
	return userreq.GetUserRequest{ID: id}, nil
}

func encodeGetUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	resp := response.(userresp.GetUserResponse)
	if resp.Error == "user not found" {
		w.WriteHeader(http.StatusNotFound)
	} else if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	return json.NewEncoder(w).Encode(resp)
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req userreq.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	resp := response.(userresp.CreateUserResponse)
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	return json.NewEncoder(w).Encode(resp)
}
