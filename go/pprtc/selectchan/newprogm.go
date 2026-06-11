package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	stop := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		close(stop) // stop after 5 seconds
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Performing scheduled task...")
		case <-stop:
			fmt.Println("Stop signal received. Exiting.")
			return
		}
	}
}

// Use Case: Periodic task execution with graceful stop.
