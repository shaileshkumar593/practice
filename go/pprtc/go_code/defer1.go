package main

import "fmt"

// defer is evaluated in LIFO order

func add(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Result :", res)
	return 0
}

func main() {
	fmt.Println("start")
	defer fmt.Println("End")
	defer add(84, 16)
	defer add(10, 10)
}
