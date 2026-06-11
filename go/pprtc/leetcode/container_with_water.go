package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {

		h := min(height[left], height[right])
		width := right - left
		area := h * width

		if area > maxWater {
			maxWater = area
		}

		// move smaller height
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	// Test cases
	testCases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7}, // expected 49
		{1, 1},                      // expected 1
		{4, 3, 2, 1, 4},             // expected 16
		{1, 2, 1},                   // expected 2
		{2, 3, 10, 5, 7, 8, 9},      // expected 36
	}

	for i, arr := range testCases {
		fmt.Printf("Test Case %d: %v → Max Water = %d\n", i+1, arr, maxArea(arr))
	}
}

/*
1. Write brute force (O(n²))
2. Optimize using two pointers (O(n))
3. Explain WHY moving smaller pointer works


“Why two pointer works?”

Answer:

We start with the maximum width and try to maximize height.
Since area is limited by the smaller height, we move the pointer
with the smaller height to potentially find a taller line.

Why min?
👉 Because water is limited by the shorter wall

|     |
|     |   ← taller wall doesn't matter
|  |  |   ← shorter wall limits water
*/
