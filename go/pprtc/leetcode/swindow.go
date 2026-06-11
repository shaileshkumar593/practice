package main

import "fmt"

func maxSumSubarray(arr []int, k int) (int, []int) {
	if len(arr) < k {
		return -1, nil
	}

	// Step 1: Initial window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	maxSum := windowSum
	startIndex := 0

	// Step 2: Slide window
	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]

		if windowSum > maxSum {
			maxSum = windowSum
			startIndex = i - k + 1
		}
	}

	// Step 3: Extract subarray
	result := arr[startIndex : startIndex+k]

	return maxSum, result
}

func main() {
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3

	sum, subarray := maxSumSubarray(arr, k)

	fmt.Println("Max Sum:", sum)       // 9
	fmt.Println("Subarray:", subarray) // [5 1 3]
}
