package main

import (
	"fmt"
)

// ✔ Single Worker — No sync, No done, No extra channel

// single worker, safe to close out
func worker4(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * n
	}
	close(out) // SAFE because only one worker writes to out
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go worker4(in, out)

	// send numbers
	go func() {
		for i := 1; i <= 10; i++ {
			in <- i
		}
		close(in)
	}()

	// read results
	for v := range out {
		fmt.Println(v)
	}
}
