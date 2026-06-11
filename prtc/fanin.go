package main

import (
	"fmt"
	"sync"
)

func FanIn(wg *sync.WaitGroup, ch1, ch2 <-chan int, out chan<- int) {

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch1 {
			out <- val
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch2 {
			out <- val * val
		}
	}()

	wg.Wait()

	close(out)

}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	out := make(chan int)
	wg := sync.WaitGroup{}

	go FanIn(&wg, ch1, ch2, out)
	go func() {
		for val := range out {
			fmt.Println(val)
		}
	}()
	for i := 0; i < 7; i++ {
		ch1 <- i
		ch2 <- i * i
	}

	close(ch1)
	close(ch2)

}
