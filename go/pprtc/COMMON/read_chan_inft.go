package main

import (
	"fmt"
	"time"
)

// producer continuously sends integers into a channel
func producer(ch chan int) {
	i := 1
	for {
		ch <- i // send data
		fmt.Printf("Produced: %d\n", i)
		i++
		time.Sleep(1 * time.Second) // simulate work
	}
}

func main() {
	ch := make(chan int)

	go producer(ch) // start producer goroutine

	// ✅ Infinite loop reading from the channel
	for {
		val := <-ch // read from channel
		fmt.Printf("Received in main: %d\n", val)
	}
}
