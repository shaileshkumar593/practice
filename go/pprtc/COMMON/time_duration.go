package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting dynamic sleep demo...")

	for i := 1; i <= 5; i++ {
		// Convert int to time.Duration (multiply by time.Second)
		sleepDuration := time.Duration(i) * time.Second

		fmt.Printf("⏳ Sleeping for %v...\n", sleepDuration)
		time.Sleep(sleepDuration) // Sleep dynamically

		fmt.Printf("✅ Woke up after %v\n", sleepDuration)
	}

	fmt.Println("🎯 All sleeps completed successfully.")
}
