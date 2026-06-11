package main

import (
	"fmt"
	"time"
)

/*
By combining select and time.After function, we can wait for a channel to become ready within a

	specified time and execute the corresponding logic if the time limit is exceeded.
*/
func main() {
	channel := make(chan int)

	select {
	case data := <-channel:
		fmt.Println("Received:", data)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout occurred.")
	}
}
