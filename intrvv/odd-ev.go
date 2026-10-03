package main

import (
	"fmt"
	"sync"
)

func OddCount(oddc chan int, evenc chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range oddc {
		fmt.Println("odd :", val)

		if val == 9 {
			// 9 is the last odd number.
			// Send 10 to even goroutine.
			evenc <- val + 1

			// Odd goroutine is finished.
			close(evenc)
			return
		}

		evenc <- val + 1
	}
}

func EvenCount(oddc chan int, evenc chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range evenc {
		fmt.Println("even :", val)

		if val == 10 {
			// 10 is the last even number.
			close(oddc)
			return
		}

		oddc <- val + 1
	}
}

func main() {
	var wg sync.WaitGroup

	oddc := make(chan int)
	evenc := make(chan int)

	wg.Add(2)

	go OddCount(oddc, evenc, &wg)
	go EvenCount(oddc, evenc, &wg)

	// Start the sequence.
	oddc <- 1

	wg.Wait()
}
