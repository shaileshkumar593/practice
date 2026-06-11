package service

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func GetUrlData(ctx context.Context, url string) (*http.Response, error) {

	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)

	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to create request: %w",
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
