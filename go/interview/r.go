package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a unit of work to be processed by a worker.
type Job struct {
	ID int
}

// Worker simulates processing of a job.
func Worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job.ID)
		time.Sleep(time.Duration(job.ID%5) * time.Second) // sleep for job.ID seconds (capped for demo)
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
	}
}

func main() {
	const numJobs = 100
	const numWorkers = 10

	jobs := make(chan Job, numJobs)
	var wg sync.WaitGroup

	// Start worker goroutines
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go Worker(w, jobs, &wg)
	}

	// Send jobs to the channel
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs) // Important: signal workers no more jobs are coming

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("✅ All jobs processed successfully!")
}
