package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go sender_msg(c)
	for i := 0; i < 9; i++ { //9 and 11 both works
		msg, ok := <-c
		if ok {
			fmt.Println(msg)
		}

	}
}

func sender_msg(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "Hello"
	}
	close(c)
}
