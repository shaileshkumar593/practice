package main

import (
	"fmt"
	"time"
)

func server1(ch chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println(<-ch)
}
func server2(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println(<-ch)
}
func main() {
	//var wg sync.WaitGroup
	input1 := make(chan int)
	input2 := make(chan int)

	go server1(input1)
	go server2(input2)

	select {
	case input1 <- 5:

	case input2 <- 7:

	}

}
