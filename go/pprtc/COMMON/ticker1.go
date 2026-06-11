package main

import (
	"fmt"
	"time"
)

func main() {
	// ✅ 1. Create a duration manually
	duration := 2 * time.Second
	fmt.Println("Duration:", duration)

	// ✅ 2. Sleep for that duration
	fmt.Printf("Sleeping for %v...\n", duration)
	time.Sleep(duration)
	fmt.Println("Woke up after sleep")

	// ✅ 3. Measure elapsed time using Duration
	start := time.Now()
	time.Sleep(1500 * time.Millisecond) // simulate work
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)

	// ✅ 4. Convert durations to different units
	fmt.Printf("Elapsed in ms: %v ms\n", elapsed.Milliseconds())
	fmt.Printf("Elapsed in ns: %v ns\n", elapsed.Nanoseconds())

	// ✅ 5. Parse duration string dynamically (e.g., from config or env)
	parsed, err := time.ParseDuration("3s")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}
	fmt.Println("Parsed duration:", parsed)

	// ✅ 6. Use duration in a ticker
	fmt.Println("Starting ticker for 1 second...")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeout := time.After(4 * time.Second) // stop after 4s
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at:", t.Format("15:04:05"))
		case <-timeout:
			fmt.Println("Timeout reached. Stopping ticker.")
			return
		}
	}
}
