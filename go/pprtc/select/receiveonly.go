package main

import (
	"fmt"
)

func printHello(p <-chan int) {
	for {
		v, ok := <-p
		if ok == false {
			fmt.Println("channel stop sending value")
			break
		}
		fmt.Println("Value received is ", v)

	}

	fmt.Println("Exiting from go routine")
}

func main() {
	c := make(chan int)

	go printHello(c)
	for k := 7; k < 15; k++ {
		c <- k * k
	}
	close(c)
	fmt.Println("Exiting from main ")
}
