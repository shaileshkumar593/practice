package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	var operation func(int, int) int
	operation = add

	fmt.Println(operation(5, 3))
}

/*

 Used in:

		Pluggable algorithms

		Configurable logic

		Dependency injection

*/
