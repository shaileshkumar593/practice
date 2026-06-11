package main

import "fmt"

func sendTomain(c chan<- int) {

	for i := 0; i < 10; i++ {
		c <- i*i + 21
	}
	close(c)
	fmt.Println("exiting from go routine")
}

func main() {
	c := make(chan int)
	go sendTomain(c)
	for {
		v, ok := <-c
		if ok == false {
			fmt.Println("channel close ")
			break // if comment then code will starve
		}
		fmt.Println("Received value ", v)
	}
	fmt.Println("Exiting from main")
}
