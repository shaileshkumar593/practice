package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	a := 0
	b := 1

	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func main() {
	fmt.Println(fibonacci(10)) // 55
}

/*
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
*/
