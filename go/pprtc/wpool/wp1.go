package main

import (
	"fmt"
)

// Simple worker pool using channels only
func workerSWP(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * job
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	// 5 workers
	for i := 1; i <= 5; i++ {
		go workerSWP(i, jobs, results)
	}

	// Sending jobs
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Receiving results
	for i := 1; i <= 10; i++ {
		fmt.Println(<-results)
	}
}
