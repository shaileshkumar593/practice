package main

import (
	"fmt"
	"time"
)

func main() {
	// Let's say we want to tick every `i` seconds dynamically
	for i := 1; i <= 3; i++ {
		interval := time.Duration(i) * time.Second // convert int to time.Duration
		fmt.Printf("\nStarting ticker for every %v...\n", interval)

		ticker := time.NewTicker(interval)
		defer ticker.Stop() // ensure cleanup

		// Stop after 3 ticks
		done := time.After(3 * interval)

		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick at:", t.Format("15:04:05"))
			case <-done:
				fmt.Printf("Ticker stopped after %v.\n", interval)
				goto NEXT // break outer loop safely
			}
		}
	NEXT:
		time.Sleep(500 * time.Millisecond) // slight pause before next interval
	}

	fmt.Println("\n✅ All dynamic tickers completed successfully.")
}
