package main

import "fmt"

func hello(p chan int) {
	fmt.Println("ok in goroutine")
	<-p
	fmt.Println("ohhh")
}

func main() {
	ch := make(chan int)
	go hello(ch)
	fmt.Println("done")

} // ok no issue output done
