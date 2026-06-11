package main

import "fmt"

func BFS(graph map[int][]int, level []int, visited map[int]bool) {
	if len(level) == 0 {
		return
	}

	levelnode := []int{}

	for _, node := range level {
		if visited[node] {
			continue
		}

		visited[node] = true
		fmt.Println(node)

		for _, neighbour := range graph[node] {
			if !visited[neighbour] {
				levelnode = append(levelnode, neighbour)
			}
		}
	}

	BFS(graph, levelnode, visited)
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

	BFS(graph, []int{1}, visited)
}
