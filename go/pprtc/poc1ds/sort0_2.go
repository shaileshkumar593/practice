package main

import (
	"fmt"
)

func main() {
	var x = []int{2, 1, 2, 1, 0, 1, 0, 2}
	fmt.Println("Array before Sorting")
	fmt.Println(x)
	sortElements(x)
	fmt.Println("Array after Sorting")
	fmt.Println(x)
}

func sortElements(x []int) {
	low := 0
	mid := 0
	high := len(x) - 1
	for mid <= high {
		if x[mid] == 0 {
			x[low], x[mid] = x[mid], x[low]
			low += 1
			mid += 1
		} else if x[mid] == 1 {
			mid += 1
		} else if x[mid] == 2 {
			x[mid], x[high] = x[high], x[mid]
			high -= 1
		}
	}
}
