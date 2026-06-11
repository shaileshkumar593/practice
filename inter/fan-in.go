package main

import (
	"fmt"
	"sync"
)

// Multiple channels → One output channel

func fanIn(ch1, ch2 <-chan int, out chan<- int) {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range ch1 {
			out <- v
		}
	}()

	go func() {
		defer wg.Done()
		for v := range ch2 {
			out <- v
		}
	}()

	wg.Wait()
	close(out)
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	out := make(chan int)

	go fanIn(ch1, ch2, out)

	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i <= 50; i += 10 {
			ch2 <- i
		}
		close(ch2)
	}()

	for v := range out {
		fmt.Println(v)
	}
}
