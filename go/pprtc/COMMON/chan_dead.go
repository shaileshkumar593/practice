package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go send_msg(c)
	for i := 0; i < 10; i++ {
		c <- "hello"
	}

}

func send_msg(c1 chan string) {
	for x := 0; x < 10; x++ {
		msg := <-c1
		fmt.Println(msg)
	}
}
