package main

import "fmt"

func hello(p chan int) {
	fmt.Println("ok in goroutine")
	fmt.Println("ohhh")
}

func main() {
	ch := make(chan int)
	go hello(ch)
	<-ch
	fmt.Println("done")

}
