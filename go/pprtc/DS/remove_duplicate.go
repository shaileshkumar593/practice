package main

import (
	"fmt"
)

func removeDuplicates(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	i := 0

	for j := 1; j < len(arr); j++ {

		if arr[j] != arr[i] {
			i++
			arr[i] = arr[j]
		}
	}

	return i + 1
}

func main() {

	arr := []int{1, 2, 2, 3, 4, 4, 4, 5, 5}

	newSize := removeDuplicates(arr)

	fmt.Println("New Size:", newSize)
	fmt.Println("Unique Elements:", arr[:newSize])
}
