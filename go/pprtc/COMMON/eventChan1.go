package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Hour)
	defer cancel()

	eventChan := make(chan string)

	// Simulate an external event
	go func() {
		time.Sleep(10 * time.Second)
		eventChan <- "Event completed"
	}()

	select {
	case msg := <-eventChan:
		fmt.Println("Received:", msg)
	case <-ctx.Done():
		fmt.Println("Timeout or cancel:", ctx.Err())
	}
}
