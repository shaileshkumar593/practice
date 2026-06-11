package main

import (
	"fmt"
	"time"
)

// worker processes jobs from the 'jobs' channel and sends results to the 'results' channel
func worker54(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d: started job %d\n", id, j)
		time.Sleep(time.Millisecond * 500) // Simulate work
		fmt.Printf("Worker %d: finished job %d\n", id, j)
		results <- j * 2 // Send result back
	}
}

func main() {
	const numJobs = 9
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 1. Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker54(w, jobs, results)
	}

	// 2. Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs will be sent

	// 3. Collect all results
	for a := 1; a <= numJobs; a++ {
		// Blocks until we have received a result for every job we sent
		<-results
	}
	// Note: We often use sync.WaitGroup here in real code to be certain all workers have exited
	// before we potentially close the results channel, but this collects all *data*.

	fmt.Println("Main: All jobs processed and results collected.")
}
