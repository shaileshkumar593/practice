package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func main() {
	var a, b int = 20, 30
	defer sum(a, b)

	fmt.Println("defer with return val")

}
