package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"go-api-project/service"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	url := "https://gorest.co.in/public/v2/users"

	resp, err := service.GetUrlData(ctx, url)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {

		http.Error(
			w,
			"failed to read response body",
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"status": "success",
			"data":   string(body),
		},
	)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	_ = ctx

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(
		map[string]string{
			"status": "healthy",
		},
	)
}
