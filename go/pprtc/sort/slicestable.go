package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4, 5, 9, 6, 8, 3, 5, 7, 99, 58, 1}
	sort.SliceStable(a, func(i, j int) bool {
		//i,j are represented for two value of the slice .
		return a[i] < a[j]
	})
	fmt.Println(a)
}

/*
 sort the slice while keeping the original order of equal elements,
 use sort.SliceStable instead.*/
