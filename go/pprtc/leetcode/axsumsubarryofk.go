package main

import "fmt"

func maxSumSubarray(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}

	windowSum := 0
	maxSum := 0

	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	maxSum = windowSum

	for i := k; i < len(nums); i++ {
		windowSum += nums[i]
		windowSum -= nums[i-k]

		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

func main() {
	nums := []int{2, 1, 5, 1, 3, 2}
	k := 3

	fmt.Println(maxSumSubarray(nums, k))
}