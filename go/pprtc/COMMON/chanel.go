package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}

		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		var i int
		i = 17
		ch <- i
		ch <- i + 18
		wg.Done()
	}(ch)
	wg.Wait()
}
