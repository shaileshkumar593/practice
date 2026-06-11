package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	i := 0
	n := len(intervals)

	// 1️⃣ Add all intervals before newInterval
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		fmt.Println(" i : ", i, " Result : ", result)
		i++
	}

	fmt.Println(" out : ", i, " Result : ", result)

	// 2️⃣ Merge overlapping intervals
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])

		fmt.Println(" in : ", i, " newInt : ", newInterval)

		i++
	}

	// Add merged interval
	result = append(result, newInterval)
	fmt.Println(" result : ", result)

	// 3️⃣ Add remaining intervals
	for i < n {
		result = append(result, intervals[i])
		fmt.Println(" i : ", i, " Result : ", result)
		i++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}

	fmt.Println(insert(intervals, newInterval))
}
