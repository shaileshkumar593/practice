package main

import (
	"fmt"
	"math"
)

func threeLargest(nums []int) (int, int, int) {
	if len(nums) < 3 {
		panic("array must have at least 3 elements")
	}

	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64

	for _, n := range nums {
		if n > first {
			third = second
			second = first
			first = n
		} else if n > second && n != first {
			third = second
			second = n
		} else if n > third && n != first && n != second {
			third = n
		}
	}

	return first, second, third
}

func main() {
	arr := []int{10, 5, 20, 8, 25, 3}
	a, b, c := threeLargest(arr)
	fmt.Println(a, b, c) // 25 20 10
}
