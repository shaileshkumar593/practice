package main

import (
	"fmt"
	"sync"
)

// printOdd prints only odd numbers.
// It receives a number from the shared channel.
// If the number is odd, it prints it and sends the next even number.
// If the number is even, it sends it back so printEven can handle it.
func printOdd(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done() // notify main goroutine when this function exits

	for {
		num := <-ch // wait until some number is available in channel

		// termination condition
		// if number becomes greater than 10,
		// pass the stop signal to the other goroutine and exit
		if num > 10 {
			ch <- num
			return
		}

		// odd goroutine should process only odd numbers
		if num%2 == 1 {
			fmt.Println(num)

			// after printing odd number,
			// send next even number to channel
			ch <- num + 1
		} else {
			// if odd goroutine receives even number,
			// it cannot process it, so it puts it back
			// for even goroutine
			ch <- num
		}
	}
}

// printEven prints only even numbers.
// It receives a number from the same shared channel.
// If the number is even, it prints it and sends the next odd number.
// If the number is odd, it sends it back so printOdd can handle it.
func printEven(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done() // notify main goroutine when this function exits

	for {
		num := <-ch // wait until some number is available in channel

		// termination condition
		// after 10, both goroutines should stop
		if num > 10 {
			ch <- num
			return
		}

		// even goroutine should process only even numbers
		if num%2 == 0 {
			fmt.Println(num)

			// after printing even number,
			// send next odd number to channel
			ch <- num + 1
		} else {
			// if even goroutine receives odd number,
			// it cannot process it, so it puts it back
			// for odd goroutine
			ch <- num
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// single unbuffered channel used for communication
	// only one number/token moves between goroutines at a time
	ch := make(chan int)

	// we have two goroutines:
	// one for odd numbers and one for even numbers
	wg.Add(2)

	go printOdd(&wg, ch)
	go printEven(&wg, ch)

	// start the sequence by sending 1
	// since 1 is odd, printOdd should process first
	ch <- 1

	// wait until both goroutines finish
	wg.Wait()
}