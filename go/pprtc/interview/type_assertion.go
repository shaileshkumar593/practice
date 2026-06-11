package main

import (
	"fmt"
)

func checkType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("String:", v)

	case int:
		fmt.Println("Int:", v)

	case float64:
		fmt.Println("Float64:", v)

	case bool:
		fmt.Println("Bool:", v)

	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	checkType("hello")
	checkType(42)
	checkType(3.14)
	checkType(true)
	checkType([]int{1, 2, 3}) // Unknown type
}
