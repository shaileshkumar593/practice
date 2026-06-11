package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	index := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[index][1] {
			intervals[index][1] = max(intervals[index][1], intervals[i][1])
		} else {
			index++
			intervals[index] = intervals[i]
		}
	}

	return intervals[:index+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Hello, 世界")

	/*intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}*/

	intervals := [][]int{
		{1, 4},
		{2, 3},
	}

	merged := merge(intervals)
	fmt.Println("merged interval")
	for _, val := range merged {
		fmt.Println(val)
	}
}
