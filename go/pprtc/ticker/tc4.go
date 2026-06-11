package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a new ticker that ticks every 500 milliseconds
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool) // Channel to signal when to stop

	go func() {
		for {
			select {
			case <-done:
				return // Exit the goroutine when done is signaled
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Simulate some work for 3 seconds
	time.Sleep(3 * time.Second)

	// Stop the ticker and signal the goroutine to exit
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped.")
}

/*
	Key characteristics of time.Ticker:

		Periodic execution: Unlike time.Timer which fires once after a specified duration, a time.Ticker continues to send signals on its channel at regular intervals.
		Channel-based communication: Tickers use a channel (ticker.C) to deliver the "tick" signal (the current time) at each interval. This integrates well with Go's concurrency model, allowing goroutines to listen for these signals.
		Initialization: A ticker is created using time.NewTicker(duration), where duration specifies the interval between ticks.
		Stopping: To release resources and prevent further ticks, a ticker should be stopped using ticker.Stop() when it is no longer needed. While recent Go versions allow the garbage collector to recover unreferenced tickers, explicitly stopping them is good practice.

*/
