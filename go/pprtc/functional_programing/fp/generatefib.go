package main

import "fmt"

// startFibonacciGenerator launches a goroutine to generate Fibonacci numbers.
func startFibonacciGenerator(out chan int, n int) {
	defer close(out)

	a, b := 0, 1
	for i := 0; i < n; i++ {
		out <- a
		a, b = b, a+b
	}
}

// fibonacci returns a channel that yields n Fibonacci numbers lazily.
func fibonacci(n int) <-chan int {
	out := make(chan int)
	go startFibonacciGenerator(out, n)
	return out
}

func main() {
	for num := range fibonacci(10) {
		fmt.Println(num)
	}
}
