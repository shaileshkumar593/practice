package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	eventChan := make(chan string)

	go func() {
		time.Sleep(20 * time.Second)
		eventChan <- fmt.Sprintln("Hello Dear")
	}()

	select {
	case msg := <-eventChan:
		fmt.Println("Received msg is ", msg)
	case <-ctx.Done():
		fmt.Println(" context timeout reached ")
	}

}

// context timeout event happens
