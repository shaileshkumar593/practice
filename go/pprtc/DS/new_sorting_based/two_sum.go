package main

import "fmt"

// AllTwoSum finds all unique pairs in an unsorted array whose sum equals target
func AllTwoSum(arr []int, target int) [][2]int {
	result := [][2]int{}
	numMap := make(map[int]int) // number -> count

	for _, num := range arr {
		complement := target - num
		if count, found := numMap[complement]; found && count > 0 {
			// Found a valid pair
			fmt.Println("count ", count)
			result = append(result, [2]int{complement, num})
			numMap[complement]-- // Decrease count to avoid duplicate pair
		} else {
			numMap[num]++
		}
	}

	return result
}

func main() {
	arr := []int{3, 5, 2, -4, 8, 11, 1, -1, 6, 7, 2}
	target := 7

	pairs := AllTwoSum(arr, target)
	if len(pairs) == 0 {
		fmt.Println("No pairs found")
	} else {
		fmt.Println("Pairs with sum", target, ":")
		for _, p := range pairs {
			fmt.Println(p[0], p[1])
		}
	}
}

/*
   for unsorted array
	Time  = O(n)
	space = O(n)
*/
