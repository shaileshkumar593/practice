package main

import (
	"fmt"
	"sync"
)

func myfunc(wg *sync.WaitGroup) {

	wg1 := sync.WaitGroup{}

	var c1 = make(chan int)

	wg1.Add(1)
	//c1 <- 55 cause deadlock due to unavailability of receiver
	go multiply(c1, &wg1)

	c1 <- 55 // works fine
	fmt.Println("multiplication is ", <-c1)
	wg1.Wait()

	wg.Done()

}

func multiply(r chan int, wg2 *sync.WaitGroup) {

	val := <-r * 55
	fmt.Println("done multiplication", val)
	r <- val
	close(r)
	wg2.Done()
	//close(r)

}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	go myfunc(&wg)

	wg.Wait()

	fmt.Println("hello from main")
}
