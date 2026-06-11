package main

import (
	"fmt"
	"sync"
	"time"
)

// performTask simulates a task for a given ID and sends a result string to the channel.
func performTask1(id int, ch chan<- string, wg *sync.WaitGroup) {
	// Note: We only defer wg.Done() if we are guaranteed to reach the end of the function.
	// A long task might time out globally, but the goroutine should still finish its work and
	// signal completion to the WaitGroup.
	defer wg.Done()

	// Simulate a task with a variable sleep time (Task 3 will take 4 seconds)
	sleepDuration := time.Duration(id+1) * time.Second
	time.Sleep(sleepDuration)

	result := fmt.Sprintf("Task %d completed in %v", id, sleepDuration)

	// Try to send the result. The receiver might have timed out and closed the channel,
	// so we use a select statement here as well to prevent this goroutine from blocking forever
	// if the main process is no longer listening.
	select {
	case ch <- result: // Send the result
		fmt.Printf("Task %d sent its result\n", id)
	case <-time.After(100 * time.Millisecond): // Small local timeout to prevent hanging
		fmt.Printf("Task %d completed but main program stopped listening (timeout occurred)\n", id)
	}
}

func main() {
	const numTasks = 5
	// Using a single UNBUFFERED channel
	taskChannel := make(chan string)
	var wg sync.WaitGroup

	// Start all task goroutines
	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		go performTask1(i, taskChannel, &wg)
	}

	// Channel to signal when all *attempted* results have been processed or timed out
	done := make(chan struct{})

	// Start a separate receiver goroutine to handle results and global timeout
	go func() {
		defer close(done) // Signal the main function that result processing is done

		// Use a global timeout of 3.5 seconds for all operations
		timeout := time.After(3500 * time.Millisecond)
		receivedCount := 0

		for receivedCount < numTasks {
			select {
			case msg := <-taskChannel:
				fmt.Println("Received:", msg)
				receivedCount++
			case <-timeout:
				fmt.Println("\nGlobal Timeout: Not all tasks finished within the time limit.")
				// We break the loop and close the 'done' channel, ending the reception process
				return
			}
		}
	}()

	// Wait for the receiver goroutine to finish its job (either by receiving all or timing out)
	<-done

	// Wait for all original task goroutines to signal completion via the WaitGroup.
	// They will all eventually finish their `time.Sleep` and call `wg.Done()`.
	wg.Wait()
	fmt.Println("\nAll goroutines have finished their execution attempts.")
}
