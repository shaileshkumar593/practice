package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min Heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func topKLargest(nums []int, k int) []int {
	h := &IntHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := []int{}
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(int))
	}
	return result
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println("Top K:", topKLargest(nums, 2))
}

/*
Maintain MinHeap of size k

Remove smallest when size exceeds k

Time Complexity (DERIVATION)
Loop runs n times
Each iteration:

Push → O(log k) (heap size ≤ k)

Sometimes Pop → O(log k)

👉 Total:

n * log k
👉 O(n log k)

🧠 Why NOT log n?
👉 Heap size is k, not n

🧠 Space Complexity
Heap stores at most k elements
👉 O(k)


*/