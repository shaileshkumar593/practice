package main

import "fmt"

func DFS1(start int, graph map[int][]int) {
	visited := make(map[int]bool)

	stack := []int{start}

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		stack = stack[:len(stack)-1]

		if visited[node] {
			continue
		}

		visited[node] = true
		fmt.Println(node)

		for _, neighbour := range graph[node] {
			if !visited[neighbour] {
				stack = append(stack, neighbour)
			}
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

	DFS1(1, graph)
}
