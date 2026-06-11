package main

import "fmt"

func flat(arr [][]int) []int {
	var res []int

	for _, sub := range arr {
		res = append(res, sub...)
	}

	return res
}

func main() {
	lst := [][]int{
		{1, 2},
		{3, 4},
	}

	result := flat(lst)

	fmt.Println(result)
}
