package main

import "fmt"

func subsets(nums []int) [][]int {

	var result [][]int

	var dfs func(index int, path []int)

	dfs = func(index int, path []int) {

		temp := append([]int{}, path...)
		result = append(result, temp)

		for i := index; i < len(nums); i++ {

			path = append(path, nums[i])

			dfs(i+1, path)

			path = path[:len(path)-1]
		}
	}

	dfs(0, []int{})

	return result
}

func main() {
	fmt.Println(subsets([]int{1, 3, 2, 4, 8, 5}))
}
