package main

import (
	"fmt"
	"sort"
)

/*
	sort.Slice internally uses Introsort (QuickSort + HeapSort + InsertionSort). We’ll simplify it to QuickSort to show index comparisons.
*/

/*
	2. Where do i and j come from?
		You do not provide i and j yourself.

		The sort package calls your function internally many times to compare elements while sorting.

		Internally, Go’s sorting algorithm (Introsort) decides which elements to compare, and for each comparison, it calls:

			less(i, j)
		with the current indices i and j.

*/
// Meeting represents a custom struct for demonstration
type Meeting struct {
	Start int
	End   int
}

func main() {
	// -----------------------------
	// 1. Integer Slice
	// -----------------------------
	ints := []int{5, 2, 8, 1}
	fmt.Println("Original ints:", ints)

	sort.Ints(ints)
	fmt.Println("Ascending ints:", ints)

	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("Descending ints:", ints)

	fmt.Println("Is sorted ascending?", sort.IntsAreSorted([]int{1, 2, 3}))
	fmt.Println()

	// -----------------------------
	// 2. Float64 Slice
	// -----------------------------
	floats := []float64{3.14, 2.71, 1.41}
	fmt.Println("Original floats:", floats)

	sort.Float64s(floats)
	fmt.Println("Ascending floats:", floats)

	sort.Sort(sort.Reverse(sort.Float64Slice(floats)))
	fmt.Println("Descending floats:", floats)

	fmt.Println("Is sorted ascending?", sort.Float64sAreSorted([]float64{1.1, 2.2, 3.3}))
	fmt.Println()

	// -----------------------------
	// 3. String Slice
	// -----------------------------
	words := []string{"banana", "apple", "cherry"}
	fmt.Println("Original words:", words)

	sort.Strings(words)
	fmt.Println("Ascending strings:", words)

	sort.Sort(sort.Reverse(sort.StringSlice(words)))
	fmt.Println("Descending strings:", words)

	fmt.Println("Is sorted ascending?", sort.StringsAreSorted([]string{"a", "b", "c"}))
	fmt.Println()

	// -----------------------------
	// 4. Custom Slice: [][]int
	// -----------------------------
	intervals := [][]int{{2, 4}, {1, 2}, {7, 8}, {5, 6}}
	fmt.Println("Original intervals:", intervals)

	// Sort by start ascending
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println("Intervals sorted by start ascending:", intervals)

	// Sort by end descending
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] > intervals[j][1]
	})
	fmt.Println("Intervals sorted by end descending:", intervals)

	// Check if sorted by start ascending
	isSorted := sort.SliceIsSorted(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println("Is intervals sorted by start ascending?", isSorted)
	fmt.Println()

	// -----------------------------
	// 5. Custom Struct Slice
	// -----------------------------
	meetings := []Meeting{
		{2, 4}, {1, 2}, {7, 8}, {5, 6},
	}
	fmt.Println("Original meetings:", meetings)

	// Sort by Start ascending
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].Start < meetings[j].Start
	})
	fmt.Println("Meetings sorted by start ascending:", meetings)

	// Sort by End descending
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].End > meetings[j].End
	})
	fmt.Println("Meetings sorted by end descending:", meetings)

	// Check if meetings sorted by Start ascending
	isSortedStruct := sort.SliceIsSorted(meetings, func(i, j int) bool {
		return meetings[i].Start < meetings[j].Start
	})
	fmt.Println("Is meetings sorted by start ascending?", isSortedStruct)
}

/*

	| Type         | Ascending                     | Descending                                          | Check if Sorted                   |
| ------------ | ----------------------------- | --------------------------------------------------- | --------------------------------- |
| `[]int`      | `sort.Ints(slice)`            | `sort.Sort(sort.Reverse(sort.IntSlice(slice)))`     | `sort.IntsAreSorted(slice)`       |
| `[]string`   | `sort.Strings(slice)`         | `sort.Sort(sort.Reverse(sort.StringSlice(slice)))`  | `sort.StringsAreSorted(slice)`    |
| `[]float64`  | `sort.Float64s(slice)`        | `sort.Sort(sort.Reverse(sort.Float64Slice(slice)))` | `sort.Float64sAreSorted(slice)`   |
| Custom `[]T` | `sort.Slice(slice, lessFunc)` | `sort.Slice(slice, reverseLessFunc)`                | `sort.SliceIsSorted(slice, less)` |
*/
