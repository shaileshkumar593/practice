package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)

	val := old[n-1]
	*h = old[:n-1]

	return val
}

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)

	heap.Push(h, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= (*h)[0] {
			heap.Pop(h)
		}

		heap.Push(h, intervals[i][1])
	}

	return h.Len()
}

func main() {
	intervals := [][]int{
		{0, 30},
		{5, 10},
		{15, 20},
	}

	rooms := minMeetingRooms(intervals)

	fmt.Println("Minimum Meeting Rooms Required:", rooms)
}

/*

	1. Sort by start.

	2. Use min heap to track end times.

	3. If current start ≥ smallest end → reuse room.

	4. Else → allocate new room.
*/
