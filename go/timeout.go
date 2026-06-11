package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		ch <- "done"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)

	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}
}
