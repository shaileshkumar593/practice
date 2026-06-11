package main

import (
	"fmt"
	"sort"
)

/*

	Goal
We want to:

Keep maximum number of non-overlapping intervals
OR
Remove minimum number of overlapping intervals

Both are the same problem.

*/

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// 1️⃣ Sort by END time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0                 // number of removed intervals
	prevEnd := intervals[0][1] // end of last kept interval

	// 2️⃣ Traverse intervals
	for i := 1; i < len(intervals); i++ {

		// Overlap condition
		if intervals[i][0] < prevEnd {
			count++ // remove this interval
		} else {
			// No overlap → keep it
			prevEnd = intervals[i][1]
		}
	}

	return count
}

func main() {
	intervals := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}

	fmt.Println(eraseOverlapIntervals(intervals)) // Output: 1
}

/*

	Remove Minimum Overlapping Intervals
	Pattern: Greedy
	Sort by: ✅ End Time (Very Important)

📌 Problem
Given intervals, remove the minimum number of intervals so the rest are non-overlapping.

Example
Input:  [[1,2],[2,3],[3,4],[1,3]]
Output: 1
👉 Remove [1,3]
Remaining intervals: [1,2],[2,3],[3,4] (No overlap)

🧠 Why Sort by END Time?
Greedy rule:

Always keep the interval that ends earliest.

Because:

It leaves maximum room for upcoming intervals.

It minimizes future conflicts.

If you sort by start time ❌
You may keep a long interval and block many small valid ones.

💡 Core Logic
After sorting by end:

Keep first interval

For each next interval:

If current.start < prev.end → overlap → remove it

Else → keep it

*/

/*

	Why “Keep the Interval That Ends Earliest”?
Think of a timeline:

0 ---- 1 ---- 2 ---- 3 ---- 4 ---- 5 ---- 6 ---- 7
Every interval occupies some space on this line.

If we choose an interval that ends late, it blocks more future space.

If we choose one that ends early, it frees more space for others.

🧠 Intuition with Example
Consider:

[1,10]
[2,3]
[4,5]
[6,7]
If you pick [1,10] first ❌
1----------------10
Now:

[2,3] overlaps

[4,5] overlaps

[6,7] overlaps

You can only keep 1 interval.

If you pick the one that ends earliest [2,3] ✅
  2--3
Then:

[4,5] → fits

[6,7] → fits

You can keep 3 intervals.

🚀 Huge difference.

💡 What Is the Greedy Strategy Actually Doing?
When we sort by end time, we are asking:

Which interval finishes first?

Because:

The earlier we finish,
The sooner we are free to accept another interval.

📌 Formal Greedy Thinking
At any step:

We have multiple intervals available.

If we choose the one with:

❌ Larger end → it occupies timeline longer

✅ Smaller end → it occupies timeline shorter

Shorter occupation = more future flexibility.

🎯 Why This Minimizes Conflicts?
Suppose we have two overlapping intervals:

A = [1,5]
B = [2,3]
Both overlap.

If we must remove one, which one should we remove?

Compare ends:

A ends at 5

B ends at 3

If we keep A:

We block everything until 5

If we keep B:

We block only until 3

Keeping B gives more opportunity later.

So we remove A.

👉 This is why when overlap happens, we remove the one with larger end.

🔬 Mathematical Intuition (Exchange Argument)
Assume:

We picked a later-ending interval instead of earlier-ending one.

We can always “exchange” it with the earlier-ending interval.

And the result will:

Never reduce solution quality

Sometimes increase it

So picking earliest-ending is always safe.

That’s why greedy works here.

🧠 Another Visualization
Imagine booking meeting rooms.

If a meeting ends at:

5 PM

or 3 PM

Which one allows more meetings later?

Obviously 3 PM.

Same logic.

🆚 Why Not Sort by Start?
Example:

[1,4]
[2,3]
Sorted by start:

[1,4]
[2,3]
If you pick first interval:
You lose the second.

But optimal solution is:
Keep [2,3]

That’s why sorting by start fails for this problem.

🧩 Core Rule Summary
When overlap happens:

If current.start < prevEnd:
    remove interval with larger end
But instead of checking both every time,
We sort by end first.

So naturally:

Earlier-ending interval comes first

We automatically keep best choice

🏆 Final One-Line Intuition
Finishing earlier gives more future freedom.

*/
