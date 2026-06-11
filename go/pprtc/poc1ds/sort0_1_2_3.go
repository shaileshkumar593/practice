package main

import "fmt"

func main() {
	//var x = []int{3, 2, 1, 0, 2, 3, 1, 0}
	var x = []int{0, 1, 2, 3, 0, 2, 3, 1, 1, 3, 0, 2}
	fmt.Println("Array before sorting")
	fmt.Println(x)
	sortelement(x)
	fmt.Println("Array after sorting")
	fmt.Println(x)
}
func sortelement(x []int) {
	low := 0
	mid := 0
	high := len(x) - 1

	for mid <= high {
		if x[mid] == 0 {
			x[low], x[mid] = x[mid], x[low]
			low += 1
			mid += 1
		} else if x[mid] == 3 {
			x[mid], x[high] = x[high], x[mid]
			high -= 1
		} else if x[mid] == 1 || x[mid] == 2 {
			mid += 1
		}
	}
	fmt.Println(x)
	for low <= high {
		if x[low] == 2 {
			x[low], x[high] = x[high], x[low]
			high -= 1
		} else {
			low += 1
		}
	}

}
