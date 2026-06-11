package main

import (
	"fmt"
)

// worker: reads from `in`, writes square to `out`
func SquareNumbers(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * val
	}
}

func main() {

	// correct slice
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	in := make(chan int)
	out := make(chan int)

	// Launch ONE worker
	go SquareNumbers(in, out)

	// Send values & read squares
	for _, v := range l {
		in <- v
		fmt.Println(<-out)
	}

	// Close channels safely
	close(in)
	close(out)
}
