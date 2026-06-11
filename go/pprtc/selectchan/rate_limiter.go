package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int)
	limiter := time.NewTicker(500 * time.Millisecond)
	defer limiter.Stop()

	go func() {
		for i := 1; i <= 10; i++ {
			requests <- i
		}
		close(requests)
	}()

	for req := range requests {
		select {
		case <-limiter.C:
			fmt.Println("Handling request", req)
		}
	}
}

//Use Case: Handle one request every 500ms — rate limiting pattern.
// This program processes incoming requests from a channel at a controlled rate using a ticker.
