// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

/*
	Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
Example 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
Example 2:
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
Example 3:
Input: intervals = [[4,7],[1,4]]
Output: [[1,7]]
Explanation: Intervals [1,4] and [4,7] are considered overlapping.


*/

func MergedInterval(intervals [][]int) [][]int {

	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		current := intervals[i]

		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			result = append(result, current)
		}
	}

	return result
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

	merged := MergedInterval(intervals)
	fmt.Println("merged interval")
	for _, val := range merged {
		fmt.Println(val)
	}
}

/*
1. why sorting dong using start time?
soln : If they are not sorted, we cannot safely decide whether the current interval overlaps with the previous merged interval.
When intervals are sorted by start time:

previous.start <= current.start
So we only need to compare:

current.start <= previous.end
That’s it.
No need to check backward or scan the entire list.

🔥 What Happens If We DON’T Sort?
Example:

[[5,7], [1,3], [2,6]]
Without sorting:

Step 1:
result = [[5,7]]

Step 2:
current = [1,3]

Check:

1 <= 7 → overlap?
Yes → but logically it shouldn’t merge this way.

Now your merged result becomes wrong because [1,3] actually comes before [5,7].

The algorithm breaks because order matters.*/

/*
q2. why we sort?
soln :
When intervals are sorted by start time:

previous.start <= current.start
So we only need to compare:

current.start <= previous.end
That’s it.
No need to check backward or scan the entire list.

🔥 What Happens If We DON’T Sort?
Example:

[[5,7], [1,3], [2,6]]
Without sorting:

Step 1:
result = [[5,7]]

Step 2:
current = [1,3]

Check:

1 <= 7 → overlap?
Yes → but logically it shouldn’t merge this way.

Now your merged result becomes wrong because [1,3] actually comes before [5,7].

The algorithm breaks because order matters.

Without sorting → you would need O(n²) comparisons.

With sorting → O(n log n) total.

If you try merging without sorting:

Start with [8,10]

Compare [1,3] → No overlap (WRONG ORDER)

Compare [2,6] → No overlap (WRONG AGAIN)

You would get:

[[8,10], [1,3], [2,6]]
*/
