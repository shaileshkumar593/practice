package main

import (
	"fmt"
	"time"
)

// performTasks simulates a task that takes 2 seconds and sends a message to the channel
func performTasks(id int, ch chan string) {
	// Simulate a task that takes 2 seconds
	time.Sleep(2 * time.Second)
	ch <- fmt.Sprintf("Task %d completed", id)
}

func main() {
	// We need a buffered channel large enough to receive messages from all 5 goroutines
	// without the main function blocking indefinitely.
	taskChannel := make(chan string, 5)

	// Launch five goroutines, giving each a unique ID
	for i := 1; i <= 5; i++ {
		go performTasks(i, taskChannel)
	}

	// This loop receives the results using the select statement for the timeout logic
	for i := 0; i < 5; i++ {
		select {
		case msg := <-taskChannel:
			fmt.Println("Received:", msg)
		case <-time.After(1 * time.Second):
			// Timeout after 1 second
			// This branch will execute *for each* task that hasn't returned yet after the first second.
			fmt.Println("Timeout: Task took too long")
		}
	}
}
