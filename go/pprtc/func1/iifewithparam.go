package main

import "fmt"

func main() {

	result := func(a, b int) int {
		return a + b
	}(2, 3)

	fmt.Println(result)
}
