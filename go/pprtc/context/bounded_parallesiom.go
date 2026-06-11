// example-8-worker-pool/main.go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker432(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d: jobs closed\n", id)
				return
			}
			fmt.Printf("worker %d: processing job %d\n", id, j)
			time.Sleep(300 * time.Millisecond) // work
		case <-ctx.Done():
			fmt.Printf("worker %d: cancelled\n", id)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan int)
	var wg sync.WaitGroup
	// start 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker432(ctx, i, jobs, &wg)
	}
	// send jobs
	go func() {
		for j := 1; j <= 10; j++ {
			jobs <- j
		}
		close(jobs)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main: canceling remaining work")
	cancel()
	wg.Wait()
	fmt.Println("main: done")
}

// Description: limit concurrency with workers that observe ctx.
// Pitfall: closing jobs vs cancel: coordinate closing so workers exit cleanly.
