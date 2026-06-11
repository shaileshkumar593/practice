package main

import (
	"fmt"
)

/*
In regular read and write operations, they will block when there is no data to read or no buffer space to write.

However, with the select statement, we can execute default logic when no data is ready, avoiding the program getting stuck in an infinite waiting state.

Non-blocking communication


*/

func main() {
	channel := make(chan int)

	select {
	case data := <-channel:
		fmt.Println("Received:", data)
	default:
		fmt.Println("No data available.")
	}
}
