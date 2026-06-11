package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func ValidateURL(rawURL string) error {
	parsedUrl, err := url.ParseRequestURI(rawURL)

	if err != nil {
		return fmt.Errorf("Invalid url")
	}

	if parsedUrl.Scheme != "http" || parsedUrl.Scheme != "https" {
		return fmt.Errorf("only http and https is allowed")
	}

	return nil
}

func GetUrlData(
	ctx context.Context,
	urlStr string,
) (*http.Response, error) {
	if err := ValidateURL(urlStr); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	// create request

	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		urlStr,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to create request : %w",
			err,
		)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf(
			"request failed: %w",
			err,
		)
	}

	return resp, nil
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	urlStr := r.URL.Query().Get("url")

	if urlStr == "" {
		http.Error(
			w,
			"url query parameter is required",
			http.StatusBadRequest,
		)

		return
	}

	resp, err := GetUrlData(ctx, urlStr)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
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

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"status": "success",
			"data":   string(body),
		},
	)
}

func main() {

	http.HandleFunc(
		"/fetch",
		UserHandler,
	)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}
