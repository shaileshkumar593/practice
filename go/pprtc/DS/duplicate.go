package main

import "fmt"

// find duplicate Element

func duplicateInArray(arr []int) []int {
	visited := make(map[int]bool, 0)
	duplicate := []int{}
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			duplicate = append(duplicate, arr[i])
		} else {
			visited[arr[i]] = true
		}
	}
	return duplicate
}

func main() {
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 2, 4, 7}))
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 3}))
}
