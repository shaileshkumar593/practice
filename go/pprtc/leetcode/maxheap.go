package main

import "fmt"

type MaxHeap struct {
	data []int
}

// Create new heap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data: []int{},
	}
}

// Get size
func (h *MaxHeap) Size() int {
	return len(h.data)
}

// Check empty
func (h *MaxHeap) IsEmpty() bool {
	return len(h.data) == 0
}

// Peek max element
func (h *MaxHeap) Peek() int {
	if h.IsEmpty() {
		return -1
	}
	return h.data[0]
}

// Insert element
func (h *MaxHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

// Extract max element
func (h *MaxHeap) ExtractMax() int {
	if h.IsEmpty() {
		return -1
	}

	max := h.data[0]
	last := h.data[len(h.data)-1]

	// Move last to root
	h.data[0] = last
	h.data = h.data[:len(h.data)-1]

	// Fix heap
	h.heapifyDown(0)

	return max
}

// Build heap from array (O(n))
func (h *MaxHeap) BuildHeap(arr []int) {
	h.data = arr

	// Start from last non-leaf node
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

// heapifyUp (bottom → top)
func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2

		if h.data[parent] >= h.data[index] {
			break
		}

		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

// heapifyDown (top → bottom)
func (h *MaxHeap) heapifyDown(index int) {
	n := len(h.data)

	for {
		left := 2*index + 1
		right := 2*index + 2
		largest := index

		if left < n && h.data[left] > h.data[largest] {
			largest = left
		}

		if right < n && h.data[right] > h.data[largest] {
			largest = right
		}

		if largest == index {
			break
		}

		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		index = largest
	}
}

// Print heap (for debugging)
func (h *MaxHeap) Print() {
	fmt.Println(h.data)
}

// ---------------- MAIN ----------------

func main() {
	h := NewMaxHeap()

	// Insert elements
	h.Insert(10)
	h.Insert(20)
	h.Insert(5)
	h.Insert(30)
	h.Insert(15)

	fmt.Print("Heap after inserts: ")
	h.Print()

	fmt.Println("Max:", h.Peek())

	fmt.Println("Extract Max:", h.ExtractMax())
	fmt.Print("Heap after extract: ")
	h.Print()

	// Build heap from array
	arr := []int{3, 9, 2, 1, 4, 5}
	h.BuildHeap(arr)

	fmt.Print("Heap after BuildHeap: ")
	h.Print()
}


/*

This is a fully working MaxHeap implementation with:

Insert  -- log(n)

ExtractMax -- log(n)

Peek   --- O(1)

BuildHeap ---- O(n)

Size / IsEmpty


parent = (i - 1) / 2
left   = 2*i + 1
right  = 2*i + 2


Priority Queue

Top K problems

Dijkstra

Scheduling systems

*/



