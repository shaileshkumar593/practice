package main

import (
	"fmt"
	"sync"
)

// Worker function: reads from readCh, processes data, writes to writeCh
func worker76(id int, readCh <-chan int, writeCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range readCh {
		// simulate processing
		result := val * 2
		fmt.Printf("Worker %d processed value %d -> %d\n", id, val, result)
		writeCh <- result
	}
}

func main() {
	const numWorkers = 10
	const numJobs = 20

	readCh := make(chan int)
	writeCh := make(chan int)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker76(i, readCh, writeCh, &wg)
	}

	// Write jobs to readCh in a separate goroutine
	go func() {
		for i := 1; i <= numJobs; i++ {
			readCh <- i
		}
		close(readCh) // Important: close read channel so workers stop
	}()

	// Read results from writeCh in a separate goroutine
	go func() {
		wg.Wait()      // Wait for all workers to finish
		close(writeCh) // Close writeCh after all results are written
	}()

	// Collect results
	for result := range writeCh {
		fmt.Println("Result received:", result)
	}
}
