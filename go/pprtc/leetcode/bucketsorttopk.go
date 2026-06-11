package main

import "fmt"

func topKFrequent(nums []int, k int) []int {

	// Step 1: Frequency map
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: Buckets
	buckets := make([][]int, len(nums)+1)

	for num, freq := range freqMap {
		buckets[freq] = append(buckets[freq], num)
	}

	// Step 3: Collect top K
	result := []int{}

	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		fmt.Println(buckets[i])
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
	nums := []int{1,1,1,2,2,3,7,8,9,7,7}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

//https://www.linkedin.com/in/shailesh-kumar-1a61ab3a/