package main

import (
	"fmt"
	"time"
)

func Receiver(ch chan int) {

	for val := range ch {
		fmt.Println(val)
	}
}
func main() {

	ch := make(chan int)

	go Receiver(ch)
	for i := 0; i < 6; i++ {
		ch <- i * i
	}

	close(ch)

	time.Sleep(4 * time.Second)
}
