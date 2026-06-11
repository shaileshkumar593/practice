package main

import (
	"fmt"
	"time"
)

func Abc(c chan int) {
	fmt.Println(<-c)
}
func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go Abc(ch)
		ch <- i
	}
	time.Sleep(5 * time.Second)
}
