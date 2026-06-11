package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(300 * time.Millisecond)
		}
		close(ch)
	}()

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed. Exiting loop.")
				return
			}
			fmt.Println("Received:", v)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Use Case: Detecting channel closure to exit loop gracefully.
// This program listens to a channel and exits the loop when the channel is closed.
