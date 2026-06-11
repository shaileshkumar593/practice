package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		data <- "Result from long task"
	}()

	for {
		select {
		case msg := <-data:
			fmt.Println("Received:", msg)
			return
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout: no response within 2s.")
			return
		}
	}
}

// Use Case: Handling slow operations with timeout fallback.
