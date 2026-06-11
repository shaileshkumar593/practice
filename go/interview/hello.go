package main

import (
	"fmt"
	"sync"
)

func Even(
	wg *sync.WaitGroup,
	ec, oc chan int,
	resultChan chan int,
) {
	defer wg.Done()

	val := 0

	resultChan <- val

	for {

		// create odd number
		val++

		// stop safely
		if val > 15 {
			close(oc)
			return
		}

		// send odd number
		oc <- val

		// receive even number
		v, ok := <-ec

		if !ok {
			close(oc)
			return
		}

		val = v

		resultChan <- val
	}
}

func Odd(
	wg *sync.WaitGroup,
	ec, oc chan int,
	resultChan chan int,
) {
	defer wg.Done()

	for {

		val, ok := <-oc

		if !ok {
			close(ec)
			return
		}

		resultChan <- val

		// stop at 15
		if val == 15 {
			close(ec)
			return
		}

		// generate even number
		val++

		ec <- val
	}
}
func main() {

	wg := sync.WaitGroup{}

	oddc := make(chan int)
	evenc := make(chan int)

	// all goroutines send results here
	resultChan := make(chan int)

	// collector owns slice
	result := []int{}

	//
	// Collector Goroutine
	//
	go func() {
		for val := range resultChan {
			result = append(result, val)
		}
	}()

	wg.Add(2)

	go Even(&wg, evenc, oddc, resultChan)
	go Odd(&wg, evenc, oddc, resultChan)

	wg.Wait()

	// ensure collector exits
	close(resultChan)

	fmt.Println(result)

	fmt.Println("Hello World")
}
