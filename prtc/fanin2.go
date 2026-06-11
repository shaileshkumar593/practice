package main

import "fmt"

func Worker(ch1, ch2 <-chan int, out chan<- int) {

	for ch1 != nil || ch2 != nil {
		select {
		case val, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			out <- val

		case val, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			out <- val
		}
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	out := make(chan int)

	go Worker(ch1, ch2, out)

	go func() {
		for i := 0; i <= 7; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i <= 7; i++ {
			ch2 <- i * i
		}
		close(ch2)
	}()

	for val := range out {
		fmt.Println(val)
	}
}
