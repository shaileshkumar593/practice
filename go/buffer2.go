package main

import "fmt"

func main() {

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Done")
}

/*
	A concise way to remember the behavior is:

close(ch) does not delete buffered values.

After a channel is closed, you can still receive all buffered values.

Once the buffer is empty, receives return the zero value and ok == false.

Sending to a closed channel causes a panic.

Receiving from a closed channel never blocks.



A good production practice is:

The goroutine that creates the channel usually owns it.

The sender is typically responsible for closing the channel.

Receivers should generally not close channels they didn't create.

A channel is a thread-safe communication mechanism used by goroutines
to exchange data and synchronize execution. It allows concurrent programs to communicate without explicitly sharing memory.

*/
