package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Worker(c <-chan int, wno int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range c {
		fmt.Println(" worker Number :", wno, " worker receive val :", val)
	}
}

func main() {
	c := make(chan int)

	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Worker(c, i, &wg)
	}

	for i := 0; i < 8; i++ {
		c <- i
	}
	close(c)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	wg.Wait()
}
