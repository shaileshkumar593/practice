package main

import (
	"fmt"
)

// worker: read from in, square, send to out
func SquareNumber2(id int, in <-chan int, out chan<- int) {
	for {
		select {
		case v, ok := <-in:
			if !ok { // input closed
				return // worker exits
			}
			out <- v * v
		}
	}
}

func main() {

	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	in := make(chan int)
	out := make(chan int)

	workerCount := 5

	// Launch 5 workers
	for i := 0; i < workerCount; i++ {
		go SquareNumber2(i, in, out)
	}

	// Producer: send numbers
	go func() {
		for _, v := range l {
			in <- v
		}
		close(in) // workers will exit on this
	}()

	// Read EXACT number of results (len(l))
	for i := 0; i < len(l); i++ {
		select {
		case res := <-out:
			fmt.Println(res)
		}
	}
}
