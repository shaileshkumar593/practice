package main

import "fmt"

func partitionInBucket(a []int, parts int) [][]int {
	n := len(a)
	result := make([][]int, parts)

	baseSize := n / parts
	extra := n % parts // remaining elements

	start := 0

	for i := 0; i < parts; i++ {
		size := baseSize

		// put all extra elements in last bucket
		if i == parts-1 {
			size += extra
		}

		end := start + size
		result[i] = a[start:end]
		start = end
	}

	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := partitionInBucket(a, 4)

	for i, part := range res {
		fmt.Printf("p%d: %v\n", i+1, part)
	}
}
