package main

import (
	"fmt"
	"sync"
	"time"
)

// worker processes jobs as they arrive over the channel
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Duration(job%5) * time.Second) // simulate variable work
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main() {
	const (
		numWorkers = 10
		numJobs    = 100
	)

	jobs := make(chan int) // ❌ no buffer
	var wg sync.WaitGroup

	// start worker goroutines
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// send jobs synchronously (blocks until a worker receives each job)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // tell workers we're done sending

	wg.Wait() // wait for all workers to finish
	fmt.Println("✅ All jobs processed successfully!")
}
