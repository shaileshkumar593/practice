package main

import "fmt"

type Operation func(int, int) int

func add(a, b int) int {
	return a + b
}

func main() {
	var op Operation = add

	fmt.Println(op(2, 3))
}
