package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{5, 3, 4, 7, 8, 9}
	sort.Slice(a, func(i, j int) bool {
		return a[j] > a[i]
	})
	fmt.Println(a) // [9 8 7 5 4 3]
}

/*
 function sort.Slice. It sorts a slice using a
 provided function less(i, j int) bool  */
