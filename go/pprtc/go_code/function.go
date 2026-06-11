package main

import "fmt"

func add(a int, b float64) float64 {
	c := float64(a) + b
	// explicit typecasting is required
	return c
}

func main() {
	a := 24
	b := 29.78
	d := add(a, b)
	fmt.Println("Sum of number is ", d)
}
