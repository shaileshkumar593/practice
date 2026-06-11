package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(<-ch)
		}()
		ch <- i
	}
	time.Sleep(5 * time.Second)
}
