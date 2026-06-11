package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		ch <- i
		go fmt.Println(<-ch)
	}
	time.Sleep(5 * time.Second)
}
