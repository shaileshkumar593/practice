package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	// Sort intervals by starting time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println("1 : ", intervals)
	result := [][]int{intervals[0]}

	fmt.Println("2 : ", result)

	for _, interval := range intervals[1:] {
		fmt.Println("3 : ", interval)
		last := result[len(result)-1]

		fmt.Println("4 : ", last)
		// If intervals overlap → merge them
		if interval[0] <= last[1] {
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		} else {
			// No overlap → add as new interval
			result = append(result, interval)
		}
	}

	fmt.Println("5 : ", result)
	return result
}

func main() {
	input := [][]int{{1, 3}, {8, 10}, {2, 6}}
	fmt.Println(merge(input))
}
