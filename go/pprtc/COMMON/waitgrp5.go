package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go dosomething(&wg)

	}

	wg.Wait()
	fmt.Println("Finish for loop")
}

func dosomething(wg *sync.WaitGroup) {
	fmt.Println("Do something")
	wg.Done()
}
