package main

import (
	"fmt"
)

func f() (x int) {
	defer func() {
		x++
	}()

	return 10
}

func main() {
	fmt.Println(f())
}
