package main

import "fmt"

// HoarePartition partitions the array using the first element as pivot
func HoarePartition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low - 1
	j := high + 1

	for {
		// Find leftmost element greater than or equal to pivot
		i++
		for arr[i] < pivot {
			i++
		}

		// Find rightmost element smaller than or equal to pivot
		j--
		for arr[j] > pivot {
			j--
		}

		// If two pointers met
		if i >= j {
			return j
		}

		// Swap arr[i] and arr[j]
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// QuickSort implements QuickSort using Hoare's partition scheme
func QuickSort(arr []int, low, high int) {
	if low < high {
		pi := HoarePartition(arr, low, high)

		// Recursively sort elements before and after partition
		QuickSort(arr, low, pi)
		QuickSort(arr, pi+1, high)
	}
}

// PrintArray prints the elements of the array
func PrintArray(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:")
	PrintArray(arr)
}
