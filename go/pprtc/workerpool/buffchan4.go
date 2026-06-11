package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)
	n, open := <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
}

/*
It's possible to read data from a already closed buffered channel.
The channel will return the data that is already written to the channel
and once all the data has been read, it will return the zero value of the channel.*/
