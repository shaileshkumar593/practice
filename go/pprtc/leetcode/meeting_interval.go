package main

import (
	"fmt"
	"sort"
)

// canAttendMeetings checks if a person can attend all meetings
func canAttendMeetings(intervals [][]int) bool {

	// Edge case: 0 or 1 meeting → always possible
	if len(intervals) <= 1 {
		return true
	}

	// Step 1️⃣: Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Step 2️⃣: Check overlap
	for i := 1; i < len(intervals); i++ {

		prevEnd := intervals[i-1][1]
		currStart := intervals[i][0]
		
		// If current meeting starts before previous ends → overlap
		if currStart <= prevEnd {
			return false
		}
	}

	return true
}

func main() {
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}

	fmt.Println(canAttendMeetings(intervals)) // false
}

/*
	Core Idea
If two meetings overlap, one person cannot attend both.

Two intervals overlap if:

current.start < previous.end
⚠ Important: We sort by start time first.

🚀 Why Sort?
Sorting ensures:

interval[i].start >= interval[i-1].start
So we only need to compare with the previous interval.

Without sorting, overlap detection is unreliable.
*/

/*
	Why Only Compare Adjacent?
Because after sorting:

start1 <= start2 <= start3 ...
If meeting i doesn’t overlap with i-1,
it can’t overlap with any earlier meeting.

This is the greedy guarantee.


 Complexity
Time: O(n log n)
Space: O(1)
*/
