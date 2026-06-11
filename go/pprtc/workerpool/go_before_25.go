package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// Schedule the call to Done() to happen when the goroutine exits.
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		// Add to the WaitGroup counter before launching the goroutine.
		wg.Add(1)
		go worker(i, &wg) // Pass WaitGroup by pointer
	}
	// Block until the counter is zero.
	wg.Wait()

	fmt.Println("All workers finished")
}
