package main

import (
	"fmt"
	"sort"
)

func main() {
	// Input slice
	numbers := []int{1, 0, 1, 0, 0, 1, 1, 0}

	fmt.Println("Before sorting:", numbers)

	// Sort ascending
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	fmt.Println("After sorting:", numbers)
}
