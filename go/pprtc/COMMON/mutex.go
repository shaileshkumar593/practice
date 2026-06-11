package main

import (
	"fmt"
	"sync"
)

type calculation struct {
	sum int
}

func main() {

	test := calculation{}
	test.sum = 0

	wg := sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go dosomething(&test, &wg)
	}
	wg.Wait()
	fmt.Println(test.sum)
}

func dosomething(test *calculation, wg *sync.WaitGroup) {
	//fmt.Println("Address of sum  = ", &test.sum) // gives 500 in every run but without this not work fine
	test.sum++
	wg.Done()
}
