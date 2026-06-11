package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// worker simulates a concurrent task
func worker(ctx context.Context, id int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	select {
	case <-time.After(time.Duration(id) * time.Second): // simulate work
		results <- fmt.Sprintf("Worker %d completed", id)
	case <-ctx.Done(): // handle cancellation
		results <- fmt.Sprintf("Worker %d canceled: %v", id, ctx.Err())
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	results := make(chan string, 5) // buffered to prevent blocking

	numWorkers := 3

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg, results)
	}

	// Close results channel once all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Listen for results or context cancellation
	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("✅ All workers finished gracefully")
}
