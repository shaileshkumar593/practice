package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter uint64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100000; i++ {
		atomic.AddUint64(&counter, 1)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()

	fmt.Println("Final Counter:", atomic.LoadUint64(&counter))
}

/*
	What this does:
		10 goroutines increment counter

		No mutex

		No race condition

		Thread-safe increment

*/
