package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello, 世界")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg1 *sync.WaitGroup) {

		fmt.Println("Shailesh")
		wg1.Done()
	}(&wg)
	wg.Wait()

}
