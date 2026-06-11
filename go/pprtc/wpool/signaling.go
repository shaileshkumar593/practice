package main

import (
	"fmt"
	"time"
)

func longRunningTask(done chan<- bool) {
	fmt.Println("Task: Started long task...")
	time.Sleep(2 * time.Second)
	fmt.Println("Task: Finishing up.")

	// Send a signal once the task is done
	done <- true
}

func main() {
	// Unbuffered channel used purely for synchronization signal
	done := make(chan bool)

	go longRunningTask(done)

	fmt.Println("Main: Waiting for task to complete...")

	// Block the main goroutine until it receives a signal from the 'done' channel
	<-done

	fmt.Println("Main: Task completed! Exiting.")
}
