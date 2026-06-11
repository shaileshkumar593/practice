package main

import "fmt"

func main() {
	// Create channel
	ch := make(chan string)

	ch <- "hello, world" // is a blocking operation.

	// Read from channel
	msg := <-ch

	// Print
	fmt.Println(msg)
}

/*
	Reason:

		Channel is unbuffered

		Write waits for reader

		But reader comes after write

		So program freezes

*/
