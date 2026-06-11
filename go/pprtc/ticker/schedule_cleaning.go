// example4_cleanup_job/main.go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Run cleanup routine
	go cleanupRoutine(ctx)

	// Let the cleanup run for 3 minutes (demo: use 5s)
	time.Sleep(5 * time.Second)

	fmt.Println("main: stopping cleanup")
	cancel()

	// Give goroutine time to print shutdown message
	time.Sleep(200 * time.Millisecond)
}

// cleanupRoutine runs every interval until context is cancelled.
func cleanupRoutine(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)
	// For demo, shorten the interval:
	// ticker := time.NewTicker(1 * time.Second)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("🧹 cleaning up expired sessions...")
			// performCleanup()
		case <-ctx.Done():
			fmt.Println("cleanup stopped:", ctx.Err())
			return
		}
	}
}

/*
	Cleanup jobs are common in servers: delete expired tokens, logs, cache, sessions.

time.NewTicker makes sure the cleanup runs periodically.

The context allows main() to stop the cleanup goroutine cleanly.*/
