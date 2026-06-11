package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- fmt.Sprintf("ch1-%d", i)
			time.Sleep(300 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for i := 1; i <= 3; i++ {
			ch2 <- fmt.Sprintf("ch2-%d", i)
			time.Sleep(500 * time.Millisecond)
		}
		close(ch2)
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println("Got:", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println("Got:", v)
		}
	}
	fmt.Println("All channels drained.")
}

// Use Case: Merge multiple input streams into one consumer.
// This program listens to two channels and processes messages from both until both are closed.
