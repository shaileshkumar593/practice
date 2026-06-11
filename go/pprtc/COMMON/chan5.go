package main

import (
	"fmt"
	"sync"
)

func myfunc(wg *sync.WaitGroup) {

	wg1 := sync.WaitGroup{}
	//var c1 = make(chan int)
	//c1 <- 55
	var x int

	wg1.Add(1)
	x = 25
	go multiply(&wg1, &x)
	wg1.Wait()

	fmt.Println("multiplication is ", x)
	wg.Done()

}

func multiply(wg *sync.WaitGroup, r *int) {

	*r = *r * 55
	fmt.Println("done multiplication")
	wg.Done()

}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	go myfunc(&wg)

	wg.Wait()

	fmt.Println("hello from main")
}
