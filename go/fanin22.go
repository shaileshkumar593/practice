package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)

		ch <- "first"
	}()

	go func() {
		time.Sleep(2 * time.Second)

		ch <- "second"
	}()

loop:
	for {

		select {
		case msg := <-ch:
			fmt.Println(msg)

		case <-time.After(3 * time.Second):

			fmt.Println("comeout")
			break loop

		}
	}
	fmt.Println("exit main")
}
