package main

import (
	"fmt"
)

func DFS(node int, graph map[int][]int, visited map[int]bool) {
	visited[node] = true
	fmt.Println(node)

	for _, neighbour := range graph[node] {
		if !visited[neighbour] {
			DFS(neighbour, graph, visited)
		}
	}
}

func main() {
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {},
		6: {},
	}

	visited := make(map[int]bool)

	DFS(1, graph, visited)
}
