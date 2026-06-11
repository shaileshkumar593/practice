package main

import (
	"fmt"
)

func findFrequencyOfArray(arr []int) {
	frequency := make(map[int]int)
	for _, item := range arr {
		if frequency[item] == 0 {
			frequency[item] = 1
		} else {
			frequency[item]++
		}
	}
	for item, count := range frequency {
		fmt.Printf("%d occurring %d times.\n", item, count)
	}
}

func main() {
	findFrequencyOfArray([]int{2, 2, 5, 1, 3, 5, 0})
}
