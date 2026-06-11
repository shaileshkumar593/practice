package main

import (
	"fmt"
	"sync"
)

func odd(oddCh, evenCh chan int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 9; i += 2 {
		<-oddCh
		out <- i // write to output channel
		evenCh <- 1
	}
}

func even(oddCh, evenCh chan int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 2; i <= 10; i += 2 {
		<-evenCh
		out <- i // write to output channel

		if i != 10 {
			oddCh <- 1
		}
	}
}

func main() {
	var wg sync.WaitGroup

	oddCh := make(chan int)
	evenCh := make(chan int)
	out := make(chan int)

	wg.Add(2)

	go odd(oddCh, evenCh, out, &wg)
	go even(oddCh, evenCh, out, &wg)

	// collector goroutine
	go func() {
		wg.Wait()
		close(out)
	}()

	oddCh <- 1

	// read from output channel
	for v := range out {
		fmt.Println(v)
	}
}
