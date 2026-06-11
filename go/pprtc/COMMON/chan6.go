package main

import (
	"fmt"
	"sync"
)

func myfunc(wg *sync.WaitGroup, c chan int) {

	res, _ := <-c
	fmt.Println("received val is ", res)
	res = res * res

	c <- res
	wg.Done()

}

func main() {
	var c1 = make(chan int)
	c1 <- 55
	wg := sync.WaitGroup{}
	wg.Add(1)

	go myfunc(&wg, c1)

	wg.Wait()

	fmt.Println("hello from main")
}
