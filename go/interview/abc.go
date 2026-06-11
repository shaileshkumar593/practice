/*You are given a list of intervals, where each interval is represented as a pair of integers:

[start, end]

with start <= end.

Your task is to merge all overlapping intervals and return the resulting list of non-overlapping intervals, sorted by their start value.





Input: [[1,3], [2,6], [8,10]]

Output: [[1,6], [8,10]]*/

package main

import (
	"fmt"
	"sort"
)

// mergeIntervals merges overlapping intervals in-place efficiently.
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// Step 1: Sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)
	merged := [][]int{intervals[0]}

	fmt.Println(merged)
	// Step 2: Merge overlapping intervals
	for _, curr := range intervals[1:] {
		last := merged[len(merged)-1]

		if curr[0] <= last[1] { // overlap
			if curr[1] > last[1] {
				last[1] = curr[1] // merge (extend end)
			}
		} else { // no overlap
			merged = append(merged, curr)
		}
	}

	return merged
}

func main() {
	arr1 := [][]int{{1, 3}, {2, 6}, {8, 10}}
	//arr2 := [][]int{{7, 8}, {1, 5}, {2, 4}, {4, 6}}

	fmt.Println("Output 1:", mergeIntervals(arr1))
	//fmt.Println("Output 2:", mergeIntervals(arr2))
}
