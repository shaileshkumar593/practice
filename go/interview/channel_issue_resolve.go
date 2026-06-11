package main

import "fmt"

func main() {
	ch := make(chan string, 1) // buffer size 1

	ch <- "hello, world" // no block

	msg := <-ch
	fmt.Println(msg)
}

/*

	Why it works?

		Because buffer size = 1
		So write does not wait for reader.

		Key Concept
Unbuffered channel:

Sender  ←→  Receiver
Must happen at same time
Buffered channel:

Sender → Buffer → Receiver
*/
