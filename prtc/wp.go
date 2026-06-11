package main

import (
	"fmt"
	"sync"
)

func Worker(wg *sync.WaitGroup, job <-chan int, id int) {
	defer wg.Done()
	for val := range job {

		fmt.Printf(" job = %d  worker = %d\n", val, id)
	}
}

func main() {
	c := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Worker(&wg, c, i)
	}

	for i := 0; i < 40; i++ {
		c <- i
	}

	close(c)
	wg.Wait()
}
