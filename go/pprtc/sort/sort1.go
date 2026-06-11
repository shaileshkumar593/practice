package main

import (
	"fmt"
	"sort"
)

func main() {

	vals := []int{3, 1, 0, 7, 2, 4, 8, 6}
	sort.Ints(vals)
	fmt.Println(vals)

	sort.Sort(sort.Reverse(sort.IntSlice(vals)))
	fmt.Println(vals)
}
