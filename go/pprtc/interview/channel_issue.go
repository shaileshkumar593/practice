package main

import "fmt"

func main() {
	// Create channel
	ch := make(chan string)

	// Write to channel (needs separate goroutine or it will deadlock)
	go func() {
		ch <- "hello, world" // is a blocking operation.
	}()

	// Read from channel
	msg := <-ch

	// Print
	fmt.Println(msg)
}

/*

is a blocking operation.

If we write and read in the same goroutine without go, it will deadlock.

*/
