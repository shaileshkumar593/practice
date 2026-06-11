package main

import (
	"fmt"
)

func removeRepeatingInts(slice []int) []int {
	n := len(slice)
	if n < 2 {
		return slice
	}

	i := 0 // index for next unique element

	for j := 0; j < n; j++ {
		duplicate := false
		// check if slice[j] already exists in slice[0..i-1]
		for k := 0; k < i; k++ {
			if slice[k] == slice[j] {
				duplicate = true
				break
			}
		}
		if !duplicate {
			slice[i] = slice[j]
			i++
		}
	}

	return slice[:i]
}

func main() {
	nums := []int{1, 2, 3, 2, 4, 1, 5, 3}
	fmt.Println("Original slice:", nums)

	result := removeRepeatingInts(nums)
	fmt.Println("After removing duplicates:", result)

	nums2 := []int{5, 5, 5, 5, 5}
	fmt.Println("\nOriginal slice:", nums2)
	result2 := removeRepeatingInts(nums2)
	fmt.Println("After removing duplicates:", result2)
}

/*
	Original slice: [1 2 3 2 4 1 5 3]
	After removing duplicates: [1 2 3 4 5]

	Original slice: [5 5 5 5 5]
	After removing duplicates: [5]

*/
