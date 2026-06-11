package main

import "fmt"

func nextGreaterElement(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := range result {
		result[i] = -1
	}

	stack := []int{}

	for i := 0; i < n; i++ {
		fmt.Println(i)
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = nums[i]
			fmt.Println(result)
		}

		stack = append(stack, i)
	}

	return result
}

func main() {
	nums := []int{2, 1, 2, 4, 3}
	fmt.Println(nextGreaterElement(nums))
}
