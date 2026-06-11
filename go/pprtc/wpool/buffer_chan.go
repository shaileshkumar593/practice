package main

import (
	"fmt"
	"sync"
	"time"
)

func performTask(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion when the goroutine finishes
	// Simulate a task that takes variable time
	time.Sleep(time.Duration(id) * time.Second)
	ch <- fmt.Sprintf("Task %d completed", id)
}

func main() {
	const numTasks = 5
	var wg sync.WaitGroup
	// Use a buffered channel for the results to prevent goroutine leaks if the channel is never read
	taskChannel := make(chan string, numTasks)

	for i := 1; i <= numTasks; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go performTask(i, taskChannel, &wg)
	}

	// Goroutine to close the channel when all tasks are done
	go func() {
		wg.Wait()          // Block until all goroutines call Done()
		close(taskChannel) // Close the channel to signal no more data will be sent
	}()

	// Loop to receive results with a timeout
	for i := 0; i < numTasks; i++ {
		select {
		case msg, ok := <-taskChannel:
			if !ok {
				fmt.Println("Channel closed, exiting loop")
				return // Exit if the channel is closed
			}
			fmt.Println("Received:", msg)
		case <-time.After(1 * time.Second): // Timeout after 3 seconds for the whole process (adjust as needed)
			fmt.Println("Timeout: Tasks took too long")
			return // Exit on timeout
		}
	}
}
