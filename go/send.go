package main

import (
	"fmt"
	"time"
)

func Receiver(c chan int) {
	p := <-c
	fmt.Println(p)

}

func main() {
	c := make(chan int)

	c <- 3
	close(c)

	go Receiver(c)

	time.Sleep(5 * time.Second)
}
