package main

import (
	"fmt"
)

func main() {
	a := []int{}
	b := []int{}
	fmt.Println(a == b)
	e := make([]int, 0)
	f := make([]int, 0)
	fmt.Println(e == f)
}
