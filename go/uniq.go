package main

import "fmt"

func findUnique(nums []int) int {

	result := 0

	for _, num := range nums {
		result ^= num
	}

	return result
}

func main() {
	nums := []int{2, 3, 2, 4, 4}

	fmt.Println(findUnique(nums))
}
