package main

import (
	"fmt"
)

func findDuplicate(arr []int) {

	for i := 0; i < len(arr); i++ {

		// Get absolute value
		index := abs(arr[i])

		// Convert to valid index
		index--

		// Already visited
		if arr[index] < 0 {

			fmt.Println("Duplicate:", index+1)
			return
		}

		// Mark visited
		arr[index] = -arr[index]
	}
}

func abs(x int) int {

	if x < 0 {
		return -x
	}

	return x
}

func main() {

	arr := []int{1,4,2,6,5,6,7}

	findDuplicate(arr)
}