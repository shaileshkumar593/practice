package main

import (
	"fmt"
	"sort"
)

func canAttend(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}

	// Sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		prevEnd := intervals[i-1][1]
		currStart := intervals[i][0]

		if prevEnd > currStart {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canAttend([][]int{{0, 30}, {5, 10}, {15, 20}})) // false
	fmt.Println(canAttend([][]int{{7, 10}, {2, 4}}))            // true
}
