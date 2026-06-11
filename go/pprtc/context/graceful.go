package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Long-running task cancelled.")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("Working on step %d...\n", i)
		}
	}
	fmt.Println("Long-running task completed.")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancel is called

	go longRunningTask(ctx)

	// Simulate some work before cancelling
	time.Sleep(2 * time.Second)

	fmt.Println("Main routine cancelling the long-running task.")
	cancel() // Trigger cancellation

	// Give time for the goroutine to react to cancellation
	time.Sleep(1 * time.Second)
}
