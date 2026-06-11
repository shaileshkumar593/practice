package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond) // Ticks every 500ms
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return // Exit the goroutine when done signal is received
			case t := <-ticker.C:
				// Perform periodic task here
				fmt.Println("Tick at", t.Format("15:04:05"))
			}
		}
	}()

	// Let the ticker run for a bit
	time.Sleep(2 * time.Second)

	// Stop the ticker and signal the goroutine to exit
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
