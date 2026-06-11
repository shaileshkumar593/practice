package main

import "fmt"

// bucket sort
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Bucket: index = frequency
	buckets := make([][]int, len(nums)+1)

	for num, freq := range freqMap {
		buckets[freq] = append(buckets[freq], num)
	}

	fmt.Println(buckets)

	result := []int{}

	// Traverse from high frequency
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}

	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 8, 3, 8, 7, 8, 7, 9}
	k := 3
	fmt.Println(topKFrequent(nums, k))
}
