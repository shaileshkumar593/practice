package main

import "fmt"

func operator(a, b int, fn func(int, int) int) int {
	return fn(a, b) // callback
}

func add(x, y int) int {
	return x + y
}

func main() {
	result := operator(10, 20, add)
	fmt.Println(result)
}
