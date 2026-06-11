package main

import (
	"fmt"
	"time"
)

// "producer" goroutine sending integers to a "consumer" (main) goroutine via a buffered channel
func producer(tasks chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producer: Sending task %d\n", i)
		tasks <- i // Send value to channel; blocks if buffer is full
		time.Sleep(100 * time.Millisecond)
	}
	close(tasks) // Close the channel to signal no more data is coming
}

func main() {
	// Create a buffered channel for integers
	const bufferSize = 3
	tasks := make(chan int, bufferSize)

	// Start the producer goroutine
	go producer(tasks)

	// Main goroutine acts as the consumer
	for task := range tasks {
		// Receiving from a channel blocks until a value is available
		fmt.Printf("Consumer: Received task %d\n", task)
	}

	fmt.Println("Main: All tasks consumed. Channel closed.")
}
