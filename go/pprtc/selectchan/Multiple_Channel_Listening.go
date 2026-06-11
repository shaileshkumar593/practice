package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Simulate async producers
	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- fmt.Sprintf("Message %d from ch1", i)
			time.Sleep(300 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for i := 1; i <= 3; i++ {
			ch2 <- fmt.Sprintf("Message %d from ch2", i)
			time.Sleep(500 * time.Millisecond)
		}
		close(ch2)
	}()

	for {
		select {
		case msg, ok := <-ch1:
			if !ok {
				ch1 = nil // disable closed channel
				continue
			}
			fmt.Println(msg)
		case msg, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println(msg)
		default:
			if ch1 == nil && ch2 == nil {
				fmt.Println("All channels closed. Exiting loop.")
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Listen to multiple producers concurrently until both finish.
