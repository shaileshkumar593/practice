package main

import (
	"fmt"
)

func main() {

	var x = []int{0, 1, 1, 0, 0, 1, 1, 0, 0}
	fmt.Println("array before sorting")
	fmt.Println(x)
	sortElements(x)
	fmt.Println("array after sorting")
	fmt.Println(x)

}
func sortElements(x []int) {
	low := 0
	high := len(x) - 1

	for low <= high {
		if x[low] == 1 {
			x[low], x[high] = x[high], x[low]
			high -= 1
		} else {
			low += 1
		}
	}
}
