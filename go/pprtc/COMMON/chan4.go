package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	for i := 0; i < 5; i++ {
		c <- "Car"
	}
	go countCat(c)
}

func countCat(c chan string) {
	for i := 0; i < 5; i++ {
		message := <-c
		fmt.Println(message)
	}
	fmt.Println("done")
}
