package main

import (
	"fmt"
)

func threeSumTarget(nums []int, target int) bool {
	n := len(nums)

	for i := 0; i < n-2; i++ {

		seen := make(map[int]bool)
		fmt.Println("1 :", seen)

		for j := i + 1; j < n; j++ {

			complement := target - nums[i] - nums[j]

			fmt.Println("2:", seen)
			if seen[complement] {
				fmt.Println("Triplet Found:", nums[i], complement, nums[j])
				//return true
			}

			seen[nums[j]] = true
			fmt.Println("3:", seen)
		}
	}

	return false
}

func main() {
	arr := []int{1, 4, 10, 2, 3, 8, 9} //{1, 4, 45, 6, 10, 8}
	target := 13

	found := threeSumTarget(arr, target)

	if !found {
		fmt.Println("No triplet found")
	}
}
