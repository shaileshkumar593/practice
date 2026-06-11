package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		c := make(chan string)
		go send_msg(c)

		msg, ok := <-c
		if ok {
			fmt.Println(msg)
		}

	}
	wg.Wait()
}

func send_msg(c chan string) {
	c <- "Hello"
	close(c)
	wg.Done()
}
