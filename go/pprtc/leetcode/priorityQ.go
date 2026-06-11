package main

import (
	"container/heap"
	"fmt"
)

// Item in priority queue
type Item struct {
	value    string
	priority int
	index    int
}

// Heap structure
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

// Min Heap
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func main() {
	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, &Item{"task1", 3, 0})
	heap.Push(pq, &Item{"task2", 1, 0})
	heap.Push(pq, &Item{"task3", 2, 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		fmt.Println(item.value, item.priority)
	}
}



/*

Smallest priority = highest priority

Use struct + heap



Insert → heap.Push
Element added at end → heapifyUp

Height of heap = log n
👉 O(log n)

Extract → heap.Pop
Swap root with last → heapifyDown
👉 Moves down height of tree → log n
👉 O(log n)

Space Complexity
Stores all elements → O(n)

Space Complexity
Stores all elements → O(n)


*/
