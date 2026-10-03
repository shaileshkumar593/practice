package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// MinHeap implements a min-heap.
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)

	value := old[n-1]
	*h = old[:n-1]

	return value
}

func topKLargest(nums []int, k int) []int {
	if k <= 0 || len(nums) == 0 {
		return []int{}
	}

	if k >= len(nums) {
		result := append([]int(nil), nums...)

		sort.Sort(sort.Reverse(sort.IntSlice(result)))

		return result
	}

	h := &MinHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)

		// Keep only K largest values.
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, h.Len())

	for i := range result {
		result[i] = heap.Pop(h).(int)
	}

	// Heap extraction gives ascending order.
	// Reverse it for descending Top K.
	sort.Sort(sort.Reverse(sort.IntSlice(result)))

	return result
}

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{10, 5, 20, 8, 15, 3, 25}, 3},
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{5, 5, 5, 5}, 2},
		{[]int{}, 3},
		{[]int{10}, 1},
		{[]int{10}, 5},
		{[]int{1, 2, 3}, 0},
		{[]int{1, 2, 3}, -1},
		{[]int{-10, -5, -20, -1}, 2},
	}

	for _, tc := range testCases {
		result := topKLargest(tc.nums, tc.k)

		fmt.Printf(
			"nums=%v k=%d top_k=%v\n",
			tc.nums,
			tc.k,
			result,
		)
	}
}
