package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	go countCat(c)

	for {
		message, ok := <-c
		if ok {
			fmt.Println(message)
		} else {
			break
		}
	}
}

func countCat(c chan string) {
	for i := 0; i < 5; i++ {
		c <- "Cat"
	}
	close(c)

}
