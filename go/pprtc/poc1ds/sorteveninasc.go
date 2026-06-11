package main

import (
	"fmt"
	"sort"
)

/**
Sort all even numbers in ascending order and then sort all
odd numbers in descending order*/

/*
Sort with custom comparator
Use the function sort.Slice. It sorts a slice using a provided function less(i, j int) bool.
To sort the slice while keeping the original order of equal elements, use sort.SliceStable instead.
*/

func main() {
	x := []int{1, 3, 2, 7, 5, 4}

	l := 0
	r := len(x) - 1
	k := 0

	for l <= r {
		for x[l]%2 != 0 {
			l += 1
			k += 1
		}
		for x[r]%2 == 0 && l <= r {
			r -= 1
		}
		if l <= r {
			x[l], x[r] = x[r], x[l]
		}
	}
	odd := x[:k]
	even := x[k:]

	fmt.Println("Odd Elements ", odd)
	fmt.Println("Even Elements ", even)
	sort.Ints(odd)
	sort.Ints(even)
	x = []int{}
	x = append(x, odd...)
	x = append(x, even...)
	fmt.Println("Elements of Array", x)

}
