package main

import "fmt"

func BFSItr(start int, graph map[int][]int) {
	visited := make(map[int]bool)

	queue := []int{start}

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]

		queue = queue[1:]

		fmt.Println(node)

		for _, neighbour := range graph[node] {
			if !visited[neighbour] {
				visited[neighbour] = true

				queue = append(queue, neighbour)

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

	BFSItr(1, graph)
}
