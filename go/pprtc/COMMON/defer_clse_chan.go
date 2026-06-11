package main

import (
	"fmt"
	"time"
)

// producer sends data and closes the channel using defer
func producer(ch chan int) {
	// ✅ Ensures channel is closed even if function exits early
	defer func() {
		close(ch)
		fmt.Println("Producer: channel closed (defer).")
	}()

	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Producer: finished sending data.")
}

func main() {
	ch := make(chan int)

	// ✅ Defer cleanup message for main
	defer fmt.Println("Main: finished processing all data.")

	go producer(ch) // start producer

	// ✅ Infinite read loop (breaks when channel is closed)
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Main: channel closed. Exiting loop.")
			break
		}
		fmt.Printf("Main: received %d\n", val)
	}
}
