package main

import (
	"fmt"
	"sync"
)

type calculation struct {
	sum   int
	mutex sync.Mutex
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
	test.mutex.Lock()
	test.sum++
	test.mutex.Unlock()
	wg.Done()
}
