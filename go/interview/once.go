package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var once sync.Once

func worker(
	ctx context.Context,
	id int,
	resultChan chan<- string,
	wg *sync.WaitGroup,
	cancel context.CancelFunc,
) {
	defer wg.Done()

	// Simulate variable work duration
	workTime := time.Duration(rand.Intn(5)+1) * time.Second

	timer := time.NewTimer(workTime)
	defer timer.Stop()

	select {

	// If cancellation happens first
	case <-ctx.Done():

		fmt.Printf("Worker %d cancelled\n", id)
		return

	// Work completed
	case <-timer.C:

		result := fmt.Sprintf(
			"Worker %d completed in %v",
			id,
			workTime,
		)

		// ONLY first goroutine can execute this
		once.Do(func() {

			fmt.Printf("Worker %d won\n", id)

			// Send first result
			resultChan <- result

			// Cancel all remaining goroutines
			cancel()
		})
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	// Root cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Buffered channel avoids blocking
	resultChan := make(chan string, 1)

	var wg sync.WaitGroup

	// Start 10 goroutines
	for i := 1; i <= 10; i++ {

		wg.Add(1)

		go worker(
			ctx,
			i,
			resultChan,
			&wg,
			cancel,
		)
	}

	// Receive ONLY first result
	result := <-resultChan

	fmt.Println("\nFirst Result Received:")
	fmt.Println(result)

	// Wait for all goroutines to exit
	wg.Wait()

	fmt.Println("\nAll goroutines cleaned up")
}