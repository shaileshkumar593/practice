package main

import (
	"container/heap"
	"fmt"
	"sort"
)

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

func topKDistinct(nums []int, k int) []int {
	if k <= 0 {
		return []int{}
	}

	unique := make(map[int]struct{})

	for _, num := range nums {
		unique[num] = struct{}{}
	}

	if len(unique) <= k {
		result := make([]int, 0, len(unique))

		for num := range unique {
			result = append(result, num)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(result)))

		return result
	}

	h := &MinHeap{}
	heap.Init(h)

	for num := range unique {
		heap.Push(h, num)

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, h.Len())

	for i := range result {
		result[i] = heap.Pop(h).(int)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(result)))

	return result
}

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{10, 10, 9, 8, 8, 7}, 3},
		{[]int{5, 5, 5}, 2},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{}, 3},
		{[]int{10}, 1},
		{[]int{10, 20}, 5},
	}

	for _, tc := range testCases {
		result := topKDistinct(tc.nums, tc.k)

		fmt.Printf(
			"nums=%v k=%d topK=%v\n",
			tc.nums,
			tc.k,
			result,
		)
	}
}
