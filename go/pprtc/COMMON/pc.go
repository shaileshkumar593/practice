package main

import (
	"fmt"
	"time"
)

// Worker function that sends data into a channel
func producer(ch chan string) {
	defer close(ch) // close channel when done sending
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Message %d", i)
		fmt.Println("Producing:", message)
		ch <- message // send data into channel
		time.Sleep(500 * time.Millisecond)
	}
}

// Consumer function that reads data from a channel
func consumer(ch chan string, done chan bool) {
	for msg := range ch { // read until channel is closed
		fmt.Println("Consumed:", msg)
	}
	done <- true // signal completion
}

func main() {
	// ✅ Declare channels
	ch := make(chan string)
	done := make(chan bool)

	// ✅ Start producer and consumer goroutines
	go producer(ch)
	go consumer(ch, done)

	// ✅ Wait for consumer to finish reading
	<-done
	fmt.Println("All messages processed.")
}
