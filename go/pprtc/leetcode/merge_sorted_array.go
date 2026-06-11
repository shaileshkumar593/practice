package main

import "fmt"

func mergeSortedArrays(a, b []int) []int {
	i, j := 0, 0
	result := []int{}

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	// Add remaining elements
	for i < len(a) {
		result = append(result, a[i])
		i++
	}

	for j < len(b) {
		result = append(result, b[j])
		j++
	}

	return result
}

func main() {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	fmt.Println(mergeSortedArrays(a, b))
}
