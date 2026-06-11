package main

import "fmt"

func factorial(n int64) int64 {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(10)) // 3628800
}
