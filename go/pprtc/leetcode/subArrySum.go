package main

import "fmt"

func subarraySum(nums []int, k int) int {
	prefixSum := 0
	count := 0

	freq := make(map[int]int)
	freq[0] = 1

	for _, num := range nums {
		prefixSum += num

		if val, exists := freq[prefixSum-k]; exists {
			count += val
		}

		freq[prefixSum]++
		fmt.Println(freq)
	}

	return count
}

func main() {
	nums := []int{10, 2, -2, -20, 10}
	k := -10

	fmt.Println(subarraySum(nums, k))
}