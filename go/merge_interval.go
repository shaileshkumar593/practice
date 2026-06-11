package main

import (
	"fmt"
	"sort"
)

func merege(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for _, currentInterval := range intervals[1:] {
		last := result[len(result)-1]

		if currentInterval[0] <= last[1] {
			if currentInterval[1] > last[1] {
				last[1] = currentInterval[1]
			}
		} else {
			result = append(result, currentInterval)
		}
	}
	return result
}

func main() {
	input := [][]int{{1, 3}, {8, 10}, {2, 6}}
	fmt.Println(merege(input))
}
