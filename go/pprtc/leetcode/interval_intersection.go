package main

import (
	"fmt"
	"sort"
)

func intervalIntersectionUnsorted(A, B [][]int) [][]int {

	// 1️⃣ Sort both interval lists
	sort.Slice(A, func(i, j int) bool {
		return A[i][0] < A[j][0]
	})

	sort.Slice(B, func(i, j int) bool {
		return B[i][0] < B[j][0]
	})

	// 2️⃣ (Optional but recommended) Merge internally
	A = mergeIntervals(A)
	B = mergeIntervals(B)

	// 3️⃣ Now apply two pointer intersection
	i, j := 0, 0
	result := [][]int{}

	for i < len(A) && j < len(B) {
		start := max(A[i][0], B[j][0])
		end := min(A[i][1], B[j][1])

		if start <= end {
			result = append(result, []int{start, end})
		}

		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}

	return result
}

// Merge overlapping intervals
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		current := intervals[i]

		if current[0] <= last[1] {
			last[1] = max(last[1], current[1])
		} else {
			result = append(result, current)
		}
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
	A := [][]int{{10, 14}, {1, 5}}
	B := [][]int{{11, 20}, {2, 6}, {8, 10}}

	fmt.Println(intervalIntersectionUnsorted(A, B))
}

/*

	Why Internal Merge Is Important
Suppose:

A = [[1,5],[3,8]]
If not merged first:

Two-pointer logic may create duplicate/incorrect intersections.

After merge:

A = [[1,8]]
*/
