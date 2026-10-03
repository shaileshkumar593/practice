package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, ok := seen[complement]; ok {
			return []int{index, i}
		}

		seen[num] = i
	}

	return []int{-1, -1}
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
		{[]int{1, 2, 3}, 10},
		{[]int{}, 5},
		{[]int{5}, 10},
		{[]int{-3, 4, 3, 90}, 0},
		{[]int{0, 0}, 0},
	}

	for _, tc := range testCases {
		result := twoSum(tc.nums, tc.target)

		fmt.Printf(
			"nums=%v target=%d result=%v\n",
			tc.nums,
			tc.target,
			result,
		)
	}
}
