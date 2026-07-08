package main

import (
	"fmt"
	"sync"
)

var counter int

var mutex sync.Mutex

func increment() {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func main() {
	fmt.Println("Hello, World!")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment()
		wg.Done()
	}
	wg.Wait()

	fmt.Println("counter :", counter)
}

// output 0
