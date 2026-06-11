package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("===================================")
	fmt.Println("1. Loop for 5 seconds using Before")
	fmt.Println("===================================")

	end := time.Now().Add(5 * time.Second)

	for time.Now().Before(end) {
		fmt.Println("Running...", time.Now().Format("15:04:05"))
		time.Sleep(1 * time.Second)
	}

	fmt.Println()

	fmt.Println("===================================")
	fmt.Println("2. Loop for 5 seconds using Since")
	fmt.Println("===================================")

	start := time.Now()

	for time.Since(start) < 5*time.Second {
		fmt.Println("Processing...", time.Now().Format("15:04:05"))
		time.Sleep(1 * time.Second)
	}

	fmt.Println()

	fmt.Println("===================================")
	fmt.Println("3. Execute every second using Ticker")
	fmt.Println("===================================")

	ticker := time.NewTicker(1 * time.Second)

	count := 0

	for range ticker.C {

		fmt.Println("Ticker:", time.Now().Format("15:04:05"))

		count++

		if count == 5 {
			ticker.Stop()
			break
		}
	}

	fmt.Println()

	fmt.Println("===================================")
	fmt.Println("4. Execute N times")
	fmt.Println("===================================")

	for i := 1; i <= 5; i++ {

		fmt.Println("Iteration:", i)

		time.Sleep(time.Second)
	}

	fmt.Println()

	fmt.Println("===================================")
	fmt.Println("5. Timeout using select")
	fmt.Println("===================================")

	ticker = time.NewTicker(1 * time.Second)
	timeout := time.After(5 * time.Second)

	for {

		select {

		case <-ticker.C:
			fmt.Println("Working...")

		case <-timeout:
			fmt.Println("Timeout Reached")
			ticker.Stop()
			goto NEXT
		}
	}

NEXT:

	fmt.Println()

	fmt.Println("===================================")
	fmt.Println("6. Cron Style Scheduler")
	fmt.Println("===================================")

	for i := 0; i < 3; i++ {

		fmt.Println("Job Executed:", time.Now().Format("15:04:05"))

		time.Sleep(2 * time.Second)
	}
}
