package main

import "fmt"

// IsSorted checks if the array is sorted ascending or descending automatically
func IsSorted(arr []int) string {
	if len(arr) < 2 {
		return "sorted"
	}

	ascending := true
	descending := true

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			ascending = false
		}
		if arr[i] < arr[i+1] {
			descending = false
		}
	}

	if ascending {
		return "ascending"
	}
	if descending {
		return "descending"
	}
	return "unsorted"
}

func main() {
	arr1 := []int{10, 20, 30, 40, 50}
	arr2 := []int{100, 90, 70, 40, 10}
	arr3 := []int{150, 80, 70, 40, 20}
	arr4 := []int{10, 50, 30, 40, 20}

	fmt.Println("arr1:", IsSorted(arr1)) // ascending
	fmt.Println("arr2:", IsSorted(arr2)) // descending
	fmt.Println("arr3:", IsSorted(arr3)) // descending
	fmt.Println("arr4:", IsSorted(arr4)) // unsorted
}
