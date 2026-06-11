package main

import (
	"fmt"
	"sort"
)

func threeSumTarget(nums []int, target int) bool {
	n := len(nums)

	// Step 1: Sort
	sort.Ints(nums)

	// Step 2: Fix one element
	for i := 0; i < n-2; i++ {

		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				fmt.Println("Triplet Found:", nums[i], nums[left], nums[right])
				return true
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return false
}

func main() {
	arr := []int{1, 4, 45, 6, 10, 8}
	target := 13

	found := threeSumTarget(arr, target)

	if !found {
		fmt.Println("No triplet found")
	}
}

/*
Time	O(n²)
Space	O(1)
*/
