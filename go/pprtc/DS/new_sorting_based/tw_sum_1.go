package main

import (
	"fmt"
	"sort"
)

// AllTwoSumSorted finds all pairs in a sorted array whose sum equals target
// Returns 1-based indices (after sorting) and their values
func AllTwoSumSorted(arr []int, target int) [][4]int {
	// Copy array to preserve original indices
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	var result [][4]int
	left, right := 0, len(sortedArr)-1

	for left < right {
		currentSum := sortedArr[left] + sortedArr[right]
		if currentSum == target {
			result = append(result, [4]int{left + 1, sortedArr[left], right + 1, sortedArr[right]})
			left++
			right--
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4, -3, 11, 0, -5, 13, 5, 6, 7, 8}
	target := 8

	pairs := AllTwoSumSorted(arr, target)
	if len(pairs) == 0 {
		fmt.Println("No pairs found")
	} else {
		fmt.Println("Pairs with sum", target, ":")
		for _, p := range pairs {
			fmt.Printf("Index1: %d, Value1: %d; Index2: %d, Value2: %d\n", p[0], p[1], p[2], p[3])
		}
	}
}

/*
	time  = O(n)
	space = O(1)
*/
