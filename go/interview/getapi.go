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

// URL Validation
func ValidateURL(rawURL string) error {

	parsedURL, err := url.ParseRequestURI(rawURL)

	if err != nil {
		return fmt.Errorf("invalid url")
	}

	// allow only http/https
	if parsedURL.Scheme != "http" &&
		parsedURL.Scheme != "https" {

		return fmt.Errorf(
			"only http and https are allowed",
		)
	}

	return nil
}

// API Call Function
func GetUrlData(
	ctx context.Context,
	urlStr string,
) (*http.Response, error) {

	// validate URL
	if err := ValidateURL(urlStr); err != nil {
		return nil, err
	}

	// timeout context
	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)

	defer cancel()

	// create request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		urlStr,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to create request: %w",
			err,
		)
	}

	// http client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// execute request
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf(
			"request failed: %w",
			err,
		)
	}

	return resp, nil
}

// API Handler
func UserHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// get url from query param
	urlStr := r.URL.Query().Get("url")

	// validate empty url
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

// MAIN
func main() {

	http.HandleFunc(
		"/fetch",
		UserHandler,
	)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}
