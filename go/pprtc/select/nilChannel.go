package main

import "fmt"

/*
Channels are a synchronization primitive in Go, used for passing messages between Go routines (goroutines).
When a channel is declared but not initialized, its value is nil.

Attempting to send or receive data on a nil channel will block indefinitely because a nil channel is neither closed nor is there another goroutine to perform the send or receive operation.
However, a nil channel serves a special purpose in a select statement and can be used to disable a particular branch in the select statement.

*/

func main() {
	var ch chan int
	fmt.Println(ch == nil) // true
}
