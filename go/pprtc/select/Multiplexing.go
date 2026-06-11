package main

import (
	"fmt"
	"time"
)

/*
One of the most common uses of the select statement is to simultaneously listen to multiple channels and perform different
 operations based on their readiness status.

 When executing the above code, the program will randomly print "one" or "two". If the channel is empty, the program will block indefinitely.
*/

// Multiplexing

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "two"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case msg := <-c2:
		fmt.Println(msg)
	}
}
