package main

import "fmt"

func twoSumSorted(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}
}

func main() {
	nums := []int{1, 2, 4, 6, 10}
	target := 8
	fmt.Println(twoSumSorted(nums, target))
}