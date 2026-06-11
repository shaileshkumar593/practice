package main

import (
	"fmt"
)

func FindOddOccurringElement(arr []int) int {
	xor := 0
	for i := 0; i < len(arr); i++ {
		xor = xor ^ arr[i]
	}
	return xor
}
func main() {
	arr := []int{1, 4, 5, 5, 1, 4, 5, 1}
	fmt.Printf("Input array is: %d\n", arr)
	fmt.Printf("Odd occurring element in given array is: %d\n", FindOddOccurringElement(arr))
}
