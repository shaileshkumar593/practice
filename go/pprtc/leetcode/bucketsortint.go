package main

import (
	"fmt"
	"sort"
)

func bucketSortInt(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// find min & max
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	bucketSize := 5
	bucketCount := (max-min)/bucketSize + 1

	buckets := make([][]int, bucketCount)

	// distribute
	for _, num := range arr {
		index := (num - min) / bucketSize
		buckets[index] = append(buckets[index], num)
	}

	// sort + merge
	result := []int{}
	for _, bucket := range buckets {
		sort.Ints(bucket)
		result = append(result, bucket...)
	}

	return result
}

func main() {
	arr := []int{42, 32, 33, 52, 37, 47, 51}
	fmt.Println(bucketSortInt(arr))
}