package main

import (
	"fmt"
	"math"
)

// get3Largest returns the three largest unique numbers from the array
func get3Largest(arr []int) []int {
	// Initialize with smallest possible values
	first, second, third := math.MinInt, math.MinInt, math.MinInt

	for _, x := range arr {
		// If current element is greater than first
		if x > first {
			third = second
			second = first
			first = x
		} else if x > second && x != first {
			third = second
			second = x
		} else if x > third && x != second && x != first {
			third = x
		}
	}

	var result []int
	if first == math.MinInt {
		return result
	}
	result = append(result, first)
	if second == math.MinInt {
		return result
	}
	result = append(result, second)
	if third == math.MinInt {
		return result
	}
	result = append(result, third)
	return result
}

func main() {
	arr := []int{12, 13, 1, 10, 34, 1}
	res := get3Largest(arr)
	if len(res) == 0 {
		fmt.Println("No elements found")
		return
	}
	for i, v := range res {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
