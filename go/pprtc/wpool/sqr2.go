package main

import (
	"fmt"
	"sync"
)

// worker reads from `in`, squares the number, writes to `out`
func SquareNumber1(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range in {
		out <- val * val
	}
}

func main() {

	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	in := make(chan int)
	out := make(chan int)

	var wg sync.WaitGroup
	workerCount := 5

	// Launch 5 workers
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go SquareNumber1(in, out, &wg)
	}

	// Close output channel after all workers are done
	go func() {
		wg.Wait()
		close(out)
	}()

	// Producer: send all numbers into `in`
	go func() {
		for _, v := range l {
			in <- v
		}
		close(in) // Important!
	}()

	// Consumer: read squared results
	for square := range out {
		fmt.Println(square)
	}

	fmt.Println("All work done")
}
