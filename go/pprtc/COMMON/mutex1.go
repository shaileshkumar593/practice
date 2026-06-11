package main

import (
	"fmt"
	"sync"
)

var sum int = 0

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go dosomething(&wg)
	}
	wg.Wait()
	fmt.Println(sum)
}

func dosomething(wg *sync.WaitGroup) {
	//fmt.Println("Address of sum = ", &sum) //gives 500 in every run but without this not work fine
	sum++
	wg.Done()
}
