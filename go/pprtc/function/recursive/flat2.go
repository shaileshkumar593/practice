package main

import (
	"fmt"
)

func flatten(input []interface{}) []int {
	stack := append([]interface{}{}, input...)

	var result []int

	for len(stack) > 0 {
		n := len(stack)

		item := stack[n-1]

		stack = stack[:n-1]

		switch val := item.(type) {
		case int:
			result = append(result, val)

		case []interface{}:
			stack = append(stack, val...)
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func main() {

	// nested input
	input := []interface{}{
		1,
		[]interface{}{
			2,
			[]interface{}{
				3,
				4,
			},
		},
		5,
	}

	result := flatten(input)

	fmt.Println("Input :", input)
	fmt.Println("Output:", result)
}
