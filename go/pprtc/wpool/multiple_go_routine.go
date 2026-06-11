package main

import (
	"fmt"
	"sync"
	"time"
)

// simulate a task that takes a variable amount of time
func performTasks54(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate variable task duration
	duration := time.Duration(id+1) * time.Second
	time.Sleep(duration)
	ch <- fmt.Sprintf("Task %d completed in %v", id, duration)
}

func main() {
	const numTasks = 5
	// Create a buffered channel to prevent blocking when sending from goroutines
	// Note: An unbuffered channel also works fine here due to the separate receiver goroutine
	taskChannel := make(chan string)
	var wg sync.WaitGroup

	// Launch five goroutines
	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		// Pass the channel and waitgroup to each goroutine
		go performTasks54(i, taskChannel, &wg)
	}

	// Start a separate goroutine to wait for all tasks and close the channel
	go func() {
		wg.Wait()
		close(taskChannel) // Close the channel once all tasks are done
	}()

	// Receive messages from the channel using a for-range loop
	// The loop terminates automatically when the channel is closed
	for msg := range taskChannel {
		fmt.Println("Received:", msg)
	}

	fmt.Println("All tasks processed.")
}
