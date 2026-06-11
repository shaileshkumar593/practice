package main

import "fmt"

func main() {

	//mul, div := mul_div(100, 7) div declare but not used

	mul, _ := mul_div(100, 7)

	fmt.Println("Multiplication of two number return ", mul)
}

func mul_div(n1 int, n2 int) (int, int) {
	return n1 * n2, n1 / n2
}
