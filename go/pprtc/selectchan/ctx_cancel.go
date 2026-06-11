package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Worker running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(5 * time.Second) // simulate main work
	fmt.Println("Main exiting.")
}

// Use Case: Cancelling goroutine work using context with timeout.
// This program demonstrates how to use context cancellation to stop a worker goroutine after a timeout.
// Graceful goroutine cancellation using context.Context.
