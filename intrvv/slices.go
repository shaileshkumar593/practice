package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {

	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	fmt.Println(slices.Equal(a, b))

	c := []string{"GO", "RUST"}
	d := []string{"go", "rust"}

	result := slices.EqualFunc(c, d, func(x, y string) bool {
		return strings.EqualFold(x, y)
	})

	fmt.Println(result)

	e := []int{1, 2, 3}
	f := []int{1, 2, 4}

	fmt.Println(slices.Compare(e, f))
	/*
		-1 → a < b
		0 → a == b
		+1 → a > b

	*/

	slices.CompareFunc(e, f, func(x, y int) int {
		return cmp.Compare(x, y)
	})

	s := []int{10, 20, 30, 20}

	fmt.Println(slices.Index(s, 20))

	s1 := []int{10, 20, 30}

	fmt.Println(slices.Contains(s1, 20))
}
