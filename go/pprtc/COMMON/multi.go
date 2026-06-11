package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go f(wg, i)
	}

	wg.Wait()
	fmt.Println("Done.")
}

func f(wg *sync.WaitGroup, n int) {
	defer wg.Done()
	fmt.Print(n, " ")
}
