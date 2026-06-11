package main

import (
	"container/heap"
	"fmt"
)

// Define heap element
type Item struct {
	num  int
	freq int
}

// Min Heap
type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	// Step 1: Frequency map
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: Min Heap
	h := &MinHeap{}
	heap.Init(h)

	for num, freq := range freqMap {
		heap.Push(h, Item{num, freq})

		// Maintain size k
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Step 3: Extract result
	result := []int{}
	for h.Len() > 0 {
		item := heap.Pop(h).(Item)
		result = append(result, item.num)
	}

	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
