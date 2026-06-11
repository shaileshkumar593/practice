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
		for low < high && x[low] == 0 {
			low += 1
		}
		for low < high && x[high] == 1 {
			high -= 1
		}
		if low <= high {
			x[low] = 0
			x[high] = 1
			low += 1
			high -= 1
		}
	}
}
