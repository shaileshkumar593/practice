package main

import "fmt"

func subsetsWithSum(arr []int, target int) [][]int {
	result := [][]int{}

	var backtrack func(int, int, []int)

	backtrack = func(index int, sum int, path []int) {
		if sum == target {
			temp := append([]int{}, path...)
			result = append(result, temp)

			// Don't return here because negative numbers
			// can make the sum come back to target.
		}

		if index == len(arr) {
			return
		}

		for i := index; i < len(arr); i++ {
			path = append(path, arr[i])

			backtrack(
				i+1,
				sum+arr[i],
				path,
			)

			// BACKTRACK
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0, []int{})

	return result
}

func main() {
	arr := []int{1, 1, 2, 3, 4, 5} // {1,1,2,4,5,-1}

	result := subsetsWithSum(arr, 4)

	for _, subset := range result {
		fmt.Println(subset)
	}
}
