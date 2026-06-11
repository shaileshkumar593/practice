package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.example.com/data", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Process response
}

/*
	we create a context with a timeout of 2 seconds using context.WithTimeout(). We then create an HTTP request with the custom context using http.NewRequestWithContext().
	 The context ensures that if the request takes longer than the specified timeout, it will be canceled. */
