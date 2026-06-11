package main

import (
	"fmt"
	"sync"
	"time"
)

// worker performs CPU or IO work
func worker555(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(20 * time.Millisecond) // simulate work
		fmt.Printf("Worker %d processed %d\n", id, job)
		results <- job * 2 // processed result
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	workerCount := 4
	var wg sync.WaitGroup

	// Fan-Out: spawn workers
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker555(i, jobs, results, &wg)
	}

	// Producer
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Fan-In: close results when all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read merged results
	for res := range results {
		fmt.Println("Result:", res)
	}
}
