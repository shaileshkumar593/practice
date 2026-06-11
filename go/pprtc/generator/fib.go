package main

import (
	"fmt"
)

// fibonacciGenerator generates a Fibonacci sequence up to a given limit
func fibonacciGenerator(limit int) chan int {
	ch := make(chan int) // Create an unbuffered channel

	go func() { // Launch a goroutine to generate numbers
		defer close(ch) // Ensure the channel is closed when the goroutine finishes

		a, b := 0, 1
		for a <= limit {
			ch <- a // Send the current Fibonacci number to the channel
			a, b = b, a+b
		}
	}()

	return ch // Return the channel
}

func main() {
	// Consume values from the generator
	for num := range fibonacciGenerator(50) {
		fmt.Println(num)
	}
}
