package main

import "fmt"

type MinHeap struct {
	data []int
}

// Create new heap
func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: []int{},
	}
}

// Size
func (h *MinHeap) Size() int {
	return len(h.data)
}

// IsEmpty
func (h *MinHeap) IsEmpty() bool {
	return len(h.data) == 0
}

// Peek (minimum element)
func (h *MinHeap) Peek() int {
	if h.IsEmpty() {
		return -1
	}
	return h.data[0]
}

// Insert element
func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

// Extract minimum element
func (h *MinHeap) ExtractMin() int {
	if h.IsEmpty() {
		return -1
	}

	min := h.data[0]
	last := h.data[len(h.data)-1]

	// Move last to root
	h.data[0] = last
	h.data = h.data[:len(h.data)-1]

	// Fix heap
	h.heapifyDown(0)

	return min
}

// Build heap from array (O(n))
func (h *MinHeap) BuildHeap(arr []int) {
	h.data = arr

	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

// heapifyUp (bottom → top)
func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2

		if h.data[parent] <= h.data[index] {
			break
		}

		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

// heapifyDown (top → bottom)
func (h *MinHeap) heapifyDown(index int) {
	n := len(h.data)

	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		}

		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest
	}
}

// Print heap
func (h *MinHeap) Print() {
	fmt.Println(h.data)
}

// ---------------- MAIN ----------------

func main() {
	h := NewMinHeap()

	// Insert elements
	h.Insert(10)
	h.Insert(5)
	h.Insert(20)
	h.Insert(2)
	h.Insert(15)

	fmt.Print("Heap after inserts: ")
	h.Print()

	fmt.Println("Min:", h.Peek())

	fmt.Println("Extract Min:", h.ExtractMin())
	fmt.Print("Heap after extract: ")
	h.Print()

	// Build heap from array
	arr := []int{9, 4, 7, 1, 3, 6, 2}
	h.BuildHeap(arr)

	fmt.Print("Heap after BuildHeap: ")
	h.Print()
}