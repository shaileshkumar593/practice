package main

import (
	"fmt"
)

func FanIn(ch1, ch2 <-chan int, out chan<- int) {

	for ch1 != nil || ch2 != nil {

		select {

		case v, ok := <-ch1:
			if !ok {
				ch1 = nil // disable this case
				continue
			}
			out <- v

		case v, ok := <-ch2:
			if !ok {
				ch2 = nil // disable this case
				continue
			}
			out <- v
		}
	}

	close(out)
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	out := make(chan int)

	// Start FanIn
	go FanIn(ch1, ch2, out)

	// Producer 1
	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// Producer 2
	go func() {
		for i := 10; i <= 50; i += 10 {
			ch2 <- i
		}
		close(ch2)
	}()

	// Consumer
	for val := range out {
		fmt.Println(val)
	}
}
