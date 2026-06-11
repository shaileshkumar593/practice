package main

import "fmt"

func generateSubarrays(arr []int) {
	n := len(arr)

	for start := 0; start < n; start++ {
		for end := start; end < n; end++ {
			fmt.Println(arr[start : end+1])
		}
	}
}

func main() {
	arr := []int{1, 2, 3}
	generateSubarrays(arr)
}
