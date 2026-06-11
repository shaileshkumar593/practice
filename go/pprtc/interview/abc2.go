package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{2, 3, 2, 5, 5, 4, 4, 7, 6, 8, 8}

	// Step 1: Sort the slice
	sort.Ints(s) // s = [2 2 3 4 4 5 5 6 7 8 8]

	// Step 2: Remove duplicates in-place
	n := len(s)
	if n == 0 {
		fmt.Println(s)
		return
	}

	j := 0 // index of last unique element
	for i := 1; i < n; i++ {
		if s[i] != s[j] {
			j++
			s[j] = s[i]
			fmt.Println(s, i, j)
		}
	}

	// Step 3: Slice to keep only unique elements
	s = s[:j+1]

	fmt.Println("Unique elements:", s)
}

/*

sort and shift element to left if not match
	[2 3 3 4 4 5 5 6 7 8 8] 2 1
	[2 3 4 4 4 5 5 6 7 8 8] 3 2
	[2 3 4 5 4 5 5 6 7 8 8] 5 3
	[2 3 4 5 6 5 5 6 7 8 8] 7 4
	[2 3 4 5 6 7 5 6 7 8 8] 8 5
	[2 3 4 5 6 7 8 6 7 8 8] 9 6
	Unique elements: [2 3 4 5 6 7 8]
*/
