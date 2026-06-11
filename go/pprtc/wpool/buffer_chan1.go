package main

import (
	"fmt"
	"sync"
	"time"
)

// performTask simulates a task for a given ID and sends a result string to the channel.
func performTask(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the WaitGroup counter when the goroutine finishes

	// Simulate a task with a variable sleep time
	sleepDuration := time.Duration(id+1) * time.Second
	time.Sleep(sleepDuration)

	result := fmt.Sprintf("Task %d completed in %v", id, sleepDuration)
	ch <- result // Send the result to the main channel
}

func main() {
	const numTasks = 5
	taskChannel := make(chan string, numTasks) // Buffered channel for results
	var wg sync.WaitGroup

	// Start multiple goroutines
	for i := 1; i <= numTasks; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go performTask(i, taskChannel, &wg)
	}

	// Use a select loop to receive results with a global timeout
	for i := 0; i < numTasks; i++ {
		select {
		case msg := <-taskChannel:
			fmt.Println("Received:", msg)
		case <-time.After(1 * time.Second): // Timeout after 3 seconds for any *individual* receive operation
			fmt.Println("Timeout: Task took too long to return a result")
		}
	}

	// Wait for all goroutines to signal completion before the main function exits
	wg.Wait()
	fmt.Println("All goroutines have finished their execution attempts.")
}
