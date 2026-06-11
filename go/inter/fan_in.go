package main

import (
	"fmt"
)

func worker(id int, input <-chan int) <-chan int {
	out := make(chan int)

	go func() {

		defer close(out)

		for val := range input {

			fmt.Printf("Worker %d processing %d\n", id, val)

			out <- val * val
		}
	}()

	return out
}

func FanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)

	for _, ch := range channels {
		go func(c <-chan int) {
			for val := range c {
				out <- val
			}
		}(ch)
	}

	return out
}

func main() {

	input := make(chan int)

	go func() {
		defer close(input)

		for i := 1; i <= 5; i++ {
			input <- i
		}
	}()

	w1 := worker(1, input)
	w2 := worker(2, input)
	w3 := worker(3, input)

	result := FanIn(w1, w2, w3)

	for i := 0; i < 5; i++ {
		fmt.Println(<-result)
	}
}
