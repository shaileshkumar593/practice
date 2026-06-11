package main

import (
	"fmt"
	"time"
)

func main() {
	eventChan := make(chan string)

	// Simulate some external event that might happen later
	go func() {
		time.Sleep(10 * time.Second) // event happens after 10s
		eventChan <- "Event occurred successfully"
	}()

	select {
	case msg := <-eventChan:
		fmt.Println(msg)
	case <-time.After(20 * time.Hour):
		fmt.Println("Timeout reached — no event occurred within 20 hours")
	}
}
