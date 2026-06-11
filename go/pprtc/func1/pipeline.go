package main

import "fmt"

func process(nums []int, callback func(int)) {

	for _, n := range nums {
		callback(n * 2)
	}
}

func main() {

	process([]int{1, 2, 3}, func(result int) {
		fmt.Println(result)
	})
}
