package main

import (
	"fmt"
	"time"
)

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println("received:", v)
		time.Sleep(2 * time.Second)
	}
}

func main() {

	ch := make(chan int, 3)

	go consumer(ch)

	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Println("sent:", i)
	}

	fmt.Println("main finished sending")

	time.Sleep(7 * time.Second)
}

/*

	Channels synchronize access to communication,
but buffered channels reduce immediate waiting,
allowing goroutines to work asynchronously.


Goroutine = concurrent execution
Buffered channel = temporary decoupling queue
Together = asynchronous communication

*/
