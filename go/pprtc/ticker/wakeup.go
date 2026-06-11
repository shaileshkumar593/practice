// example5_worker_signal/main.go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker(ctx)

	// Run worker for 1 second
	time.Sleep(1 * time.Second)
	fmt.Println("main: stopping worker")
	cancel()

	time.Sleep(200 * time.Millisecond)
}

// worker wakes up every interval to process small batches of work.
func worker(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("👷 worker processing 1 batch...")
			// processBatch()
		case <-ctx.Done():
			fmt.Println("worker stopped")
			return
		}
	}
}

/*
	This is used for:

	message brokers

	rate-limited background jobs

	periodic polling

	batching systems

Workers wait for the ticker, process a batch, then sleep again.
*/
