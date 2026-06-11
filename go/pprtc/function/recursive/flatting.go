package main

import (
	"fmt"
)

func falletn(arr []any) []int {
	var result []int

	for _, v := range arr {
		switch x := v.(type) {
		case int:
			result = append(result, x)

		case []interface{}:
			result = append(result, falletn(x)...)

		}
	}

	return result
}

func main() {
	arr := []any{1,
		[]any{2,
			[]any{3,
				4},
		}}

	result := falletn(arr)

	fmt.Println(result)
}
