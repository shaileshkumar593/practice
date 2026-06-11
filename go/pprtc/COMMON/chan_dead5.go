package main

import (
	"fmt"
)

func message(c chan string) {
	for i := 1; i <= 10; i++ {
		c <- "Hello"
	}
	close(c)
}

func main() {
	c := make(chan string)
	go message(c)
	for i := 0; i <= 10; i++ {
		fmt.Println(<-c)
	}
}
