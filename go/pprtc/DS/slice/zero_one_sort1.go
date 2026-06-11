package main

import "fmt"

func main() {
	numbers := []int{1, 0, 1, 0, 0, 1, 1, 0}

	fmt.Println("Before:", numbers)

	left, right := 0, len(numbers)-1

	for left < right {
		// Move left index forward if already 0
		if numbers[left] == 0 {
			left++
			continue
		}

		// Move right index backward if already 1
		if numbers[right] == 1 {
			right--
			continue
		}

		// Swap when left has 1 and right has 0
		numbers[left], numbers[right] = numbers[right], numbers[left]
		left++
		right--
	}

	fmt.Println("After :", numbers)
}
