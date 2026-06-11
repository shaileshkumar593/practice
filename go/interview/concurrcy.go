package main

import (
	"fmt"
	"sync"
)

var counter int

func Increment(wg *sync.WaitGroup, jobs chan int) {
	defer wg.Done()
	for val := range jobs {
		fmt.Println(val * val)
	}
}

func main() {

	var wg sync.WaitGroup
	var jobs = make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Increment(&wg, jobs)
	}

	for j := 0; j < 1000; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	println("Final Counter:", counter)
}
