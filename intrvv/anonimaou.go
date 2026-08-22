package main

import "fmt"

func calculate(a, b int, operation func(int, int) int) int {
	result := operation(a, b)
	return result
}
func main() {
	val := calculate(20, 15, func(x, y int) int {
		return x + y
	})

	fmt.Println(val)
}
