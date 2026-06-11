package main

import (
	"fmt"
	"time"
)

// writeChannel writes data into the channel
func writeChannel(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("Written:", i)
		time.Sleep(300 * time.Millisecond)
	}
	close(ch) // close after writing
}

// readChannel reads data from the channel
func readChannel(ch <-chan int) {
	for val := range ch {
		fmt.Println("Read:", val)
	}
	fmt.Println("Channel closed, reader exiting.")
}

func main() {
	ch := make(chan int)

	// start writer
	go writeChannel(ch)

	// start reader
	readChannel(ch)

	//time.Sleep(5 * time.Second)
}
