package main

import (
	"fmt"
	"time"
)

// producer sends 5 integers, then closes the channel
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch) // close the channel after sending all data
	fmt.Println("Producer closed the channel.")
}

func main() {
	ch := make(chan int)

	go producer(ch) // start producer goroutine

	// ✅ Infinite loop that stops when channel is closed
	for {
		val, ok := <-ch // receive value and check if channel is open
		if !ok {
			fmt.Println("Channel closed. Exiting loop.")
			break
		}
		fmt.Printf("Received in main: %d\n", val)
	}

	fmt.Println("All integers processed successfully.")
}
