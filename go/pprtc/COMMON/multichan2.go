package main

import (
	"fmt"
	"time"
)

func Operation1(ch chan string, msg string) {
	time.Sleep(2 * time.Second)
	ch <- fmt.Sprintln("Hello ", msg)
}

func Operation2(ch chan string, msg string) {
	time.Sleep(4 * time.Second)
	ch <- fmt.Sprintf("Hello ", msg)
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go Operation1(chan1, "Shailesh")
	go Operation2(chan2, "Ganesh")
	// randomly select channel is time is same i.e both avail at same time
	select {
	case op1 := <-chan1:
		fmt.Println(" operation on chan1 is active ", op1)

	case op2 := <-chan2:
		fmt.Println("Operation on chan2 is active ", op2)

	default:
		fmt.Println("no chhannel is active ")
	}
}
