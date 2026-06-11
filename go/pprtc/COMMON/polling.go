package main

import (
	"fmt"
	"time"
)

// Simulated queue (in-memory slice)
var queue = []string{
	"Job-A",
	"Job-B",
	"Job-C",
}

// pollQueue checks for new data every second
func pollQueue() {
	ticker := time.NewTicker(1 * time.Second) // poll every 1 second
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if len(queue) > 0 {
				// Pop first element (FIFO)
				job := queue[0]
				queue = queue[1:]

				fmt.Printf("🟢 Processing %s at %s\n", job, time.Now().Format("15:04:05"))
				time.Sleep(500 * time.Millisecond) // simulate work
			} else {
				fmt.Printf("🕐 No new jobs at %s\n", time.Now().Format("15:04:05"))
			}

			// Optional: add stop condition if needed
			// case <-ctx.Done():
			//     fmt.Println("Stopping polling.")
			//     return
		}
	}
}

func main() {
	fmt.Println("🚀 Queue Polling Started...")
	go pollQueue()

	// Simulate adding new jobs dynamically
	time.Sleep(4 * time.Second)
	queue = append(queue, "Job-D", "Job-E")
	fmt.Println("📦 Added Job-D and Job-E")

	// Let polling run for a while
	time.Sleep(8 * time.Second)
	fmt.Println("🛑 Main finished.")
}
