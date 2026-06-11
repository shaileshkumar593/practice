package main

import (
	"fmt"
	"time"
)

func stream(ch chan string) {
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("event-%d", i)
		time.Sleep(time.Second)
	}
	close(ch)
}

func main() {
	ch := make(chan string)

	go stream(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}

//  contineous prducing and consuming
