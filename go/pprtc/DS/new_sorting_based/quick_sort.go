package main

import "fmt"

// quickSort sorts arr[left:right] in-place
func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	pivotIndex := partition(arr, left, right)
	quickSort(arr, left, pivotIndex-1)
	quickSort(arr, pivotIndex+1, right)
}

// partition partitions the array around a pivot and returns the pivot index
func partition(arr []int, left, right int) int {
	pivot := arr[right] // choose last element as pivot
	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	fmt.Println("i+	1", i+1)
	return i + 1
}

func main() {
	arr := []int{1, -1, 55, 0, -77, 88, -3, -2, 7, 5, 11, 6}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}
