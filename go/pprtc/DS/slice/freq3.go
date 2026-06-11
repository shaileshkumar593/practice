package main

import "fmt"

func numberFrequency(nums []int) map[int]int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	return freq
}

func main() {
	numbers := []int{1, 2, 3, 2, 3, 3, 4, 5, 1}
	freq := numberFrequency(numbers)

	for num, count := range freq {
		fmt.Printf("%d: %d\n", num, count)
	}
}
