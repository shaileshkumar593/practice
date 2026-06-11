package main

import "fmt"

func Multiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func main() {
	double := Multiplier(2)
	fmt.Println(double(3))
}
