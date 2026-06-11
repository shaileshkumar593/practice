package main

import (
	"fmt"
)

func main() {
	a := []int{1, 4, 7, 2, 2}
	visited := make(map[int]bool)
	for _, val := range a {
		if visited[val] == true {
			fmt.Println("Visited  is ", val)
		} else {
			visited[val] = true
		}
	}
}
