package main

import (
	"fmt"
)

func Counter(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
func main() {
	var ch = make(chan int)

	go Counter(ch)
	for i := 0; i < 10; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
	fmt.Println("End of main")
}
