package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, weight int
}

type Node struct {
	vertex int
	dist   int
}

type MinHeap []Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	node := old[n-1]
	*h = old[:n-1]
	return node
}

func dijkstra(graph map[int][]Edge, start int) map[int]int {
	dist := make(map[int]int)

	for node := range graph {
		dist[node] = math.MaxInt32
	}
	dist[start] = 0

	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Node{start, 0})

	for h.Len() > 0 {
		curr := heap.Pop(h).(Node)

		if curr.dist > dist[curr.vertex] {
			continue
		}

		for _, edge := range graph[curr.vertex] {
			newDist := curr.dist + edge.weight
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				heap.Push(h, Node{edge.to, newDist})
			}
		}
	}

	return dist
}

func main() {
	graph := map[int][]Edge{
		0: {{1, 4}, {2, 1}},
		1: {{3, 1}},
		2: {{1, 2}, {3, 5}},
		3: {},
	}

	dist := dijkstra(graph, 0)
	fmt.Println("Shortest distances:", dist)
}

/*
Dijkstra Algorithm (Min Heap)
🧠 Idea
Always pick node with smallest distance


Time Complexity (DETAILED)
Let:

V = vertices

E = edges

Key operations:
Each node can be pushed multiple times
👉 In worst case → O(E)

Each push/pop → O(log V)

Total:
E * log V
👉 O(E log V)

🧠 Why log V?
Heap size ≤ V (nodes)

Time Complexity
Insert all tasks:
n * log n
Extract all:
n * log n
👉 Total:

O(n log n)
🧠 Space Complexity
Heap stores all tasks → O(n)

🔥 FINAL COMPARISON (VERY IMPORTANT)
Problem	Time Complexity	Space
Priority Queue	O(log n) per op	O(n)
Top K	O(n log k)	O(k)
Dijkstra	O(E log V)	O(V+E)
Scheduling	O(n log n)	O(n)
*/
