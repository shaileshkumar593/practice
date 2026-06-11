package main

import (
	"fmt"
	"time"
)

// Worker Processing Jobs with Stop Signal
// Use Case: Graceful worker shutdown with stop channel (chan struct{}).

func worker(jobs <-chan int, stop <-chan struct{}) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("Job channel closed, worker exiting.")
				return
			}
			fmt.Printf("Processing job %d...\n", job)
			time.Sleep(500 * time.Millisecond)
		case <-stop:
			fmt.Println("Stop signal received. Exiting worker.")
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	jobs := make(chan int)
	stop := make(chan struct{})

	go worker(jobs, stop)

	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	time.Sleep(2 * time.Second)
	close(stop) // broadcast stop signal
	time.Sleep(time.Second)
}
