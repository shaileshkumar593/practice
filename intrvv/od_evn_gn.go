package main

import (
	"fmt"
	"sync"
)

func OddCount(
	oddc chan int,
	evenc chan int,
	max int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for val := range oddc {
		fmt.Println("odd :", val)

		// If this is the last value, don't generate another value.
		if val >= max {
			close(evenc)
			return
		}

		evenc <- val + 1
	}
}

func EvenCount(
	oddc chan int,
	evenc chan int,
	max int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for val := range evenc {
		fmt.Println("even :", val)

		// If this is the last value, don't generate another value.
		if val >= max {
			close(oddc)
			return
		}

		oddc <- val + 1
	}
}

func main() {
	var wg sync.WaitGroup

	var max int
	fmt.Print("Read last limit of generation :")
	fmt.Scan(&max)

	if max < 1 {
		fmt.Println("max must be greater than 0")
		return
	}
	fmt.Println("\n")
	oddc := make(chan int)
	evenc := make(chan int)

	wg.Add(2)

	go OddCount(oddc, evenc, max, &wg)
	go EvenCount(oddc, evenc, max, &wg)

	// Start sequence.
	oddc <- 1

	wg.Wait()
}
