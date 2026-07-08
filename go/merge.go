package main

import (
	"fmt"
)

func merge(a1 []int, b1 []int) []int {
	result := []int{}

	i, j := 0, 0
	for i < len(a1) && j < len(b1) {
		if a1[i] <= b1[j] {
			result = append(result, a1[i])
			i++
		} else if a1[i] >= b1[j] {
			result = append(result, b1[j])
			j++
		}
	}

	if i < len(a1) {
		result = append(result, a1[i:]...)
	}
	if j < len(b1) {
		result = append(result, b1[j:]...)
	}

	return result
}
func main() {
	fmt.Println("Hello, World!")

	a1 := []int{2, 6, 9, 12, 15}
	a2 := []int{1, 5, 8, 10, 11, 16, 22}

	l := merge(a1, a2)
	fmt.Println(l) // println donot print complex type. therefore it prints only address

}
