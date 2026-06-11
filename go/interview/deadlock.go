package main

import (
	"fmt"
)

func Sender(c chan int) {

	for i := 0; i <= 5; i++ {
		c <- i
	}

}

func Receiver(c chan int) {

	for j := 0; j < 5; j++ {
		p := <-c

		fmt.Println(p)
	}
}

func main() {
	fmt.Println("Hello, World!")
	c := make(chan int)
	go Sender(c)
	go Receiver(c)

	select {} // deadlock due to select
}
