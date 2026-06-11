package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("i am in goroutine")

	}(&wg)
	wg.Wait()

	fmt.Println("i am in main")
}
