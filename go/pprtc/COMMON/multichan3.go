package main

import (
	"fmt"
	"time"
)

func Operation1(ch chan string, msg string) {
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintln(" Hello from First Channel ", msg)
}

func Operation2(ch chan string, msg string) {
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintln("Hello from second chan ", msg)
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go Operation1(chan1, "Shailesh")
	go Operation2(chan2, "Vinod")

	select {
	case op1 := <-chan1:
		fmt.Println(" Channel first is Active receive Message ", op1)

	case op2 := <-chan2:
		fmt.Println("Channel second is active and received mesage ", op2)
	}
}

/*
	PS C:\Users\shailesh.kumar\goconcurrency\COMMON> go run .\multichan3.go
 Channel first is Active receive Message   Hello from First Channel  Shailesh

PS C:\Users\shailesh.kumar\goconcurrency\COMMON> go run .\multichan3.go
 Channel first is Active receive Message   Hello from First Channel  Shailesh

PS C:\Users\shailesh.kumar\goconcurrency\COMMON> go run .\multichan3.go
Channel second is active and received mesage  Hello from second chan  Vinod

	Random selection of channel
*/
