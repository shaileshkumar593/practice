package main

import "fmt"

func MyFunc(mychannel chan int) {
	for v := 0; v < 20; v++ {
		mychannel <- v
	}

	close(mychannel)
}

func main() {
	c := make(chan int)

	go MyFunc(c)
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("channel is close ", ok)
			break
		} else {
			fmt.Println("channel sent value ", res)
		}
	}
}
