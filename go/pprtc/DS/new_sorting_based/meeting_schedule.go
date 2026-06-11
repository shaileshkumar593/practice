package main

/*
Check whether all meetings can be attended without time conflicts by first sorting the meeting intervals
based on their start times. After sorting, it compares each meeting's end time with the next meeting's start
time to detect any overlaps. If an overlap is found, it returns false, indicating a scheduling conflict. Otherwise,
it returns true, meaning all meetings can be attended without overlaps.*/
import (
	"fmt"
	"sort"
)

// canAttend checks if a person can attend all meetings without overlap
func canAttend(intervals [][]int) bool {
	n := len(intervals)
	if n == 0 {
		return true
	}

	// Sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Check for overlaps
	for i := 0; i < n-1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}

func main() {
	arr := [][]int{
		{2, 4},
		{1, 2},
		{7, 8},
		{5, 6},
		{6, 8},
	}

	if canAttend(arr) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
