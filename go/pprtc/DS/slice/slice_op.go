package main

import (
	"fmt"
	"sort"
)

// Helper function to print slices
func printSlice[T any](label string, s []T) {
	fmt.Printf("%s: %v (len=%d cap=%d)\n", label, s, len(s), cap(s))
}

func main() {
	// ------------------- Slice Declaration -------------------
	var s1 []int               // nil slice
	s2 := make([]int, 5)       // slice with length 5, default values 0
	s3 := make([]int, 3, 10)   // length 3, capacity 10
	s4 := []int{1, 2, 3, 4, 5} // slice literal

	printSlice("s1 (nil)", s1)
	printSlice("s2", s2)
	printSlice("s3", s3)
	printSlice("s4", s4)

	// ------------------- Append -------------------
	s4 = append(s4, 6, 7)
	printSlice("s4 after append", s4)

	// ------------------- Subslice -------------------
	sub := s4[1:4] // elements at index 1,2,3
	printSlice("subslice s4[1:4]", sub)

	// ------------------- Copy -------------------
	dst := make([]int, len(s4))
	n := copy(dst, s4)
	fmt.Printf("Copied %d elements: %v\n", n, dst)

	// ------------------- Delete element -------------------
	index := 2
	s4 = append(s4[:index], s4[index+1:]...)
	printSlice("s4 after deleting index 2", s4)

	// ------------------- Reverse slice -------------------
	for i, j := 0, len(s4)-1; i < j; i, j = i+1, j-1 {
		s4[i], s4[j] = s4[j], s4[i]
	}
	printSlice("s4 reversed", s4)

	// ------------------- Sum and Average -------------------
	sum := 0
	for _, v := range s4 {
		sum += v
	}
	avg := float64(sum) / float64(len(s4))
	fmt.Printf("Sum: %d, Average: %.2f\n", sum, avg)

	// ------------------- Max / Min -------------------
	max, min := s4[0], s4[0]
	for _, v := range s4 {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Printf("Max: %d, Min: %d\n", max, min)

	// ------------------- Check if sorted -------------------
	isSorted := true
	for i := 1; i < len(s4); i++ {
		if s4[i-1] > s4[i] {
			isSorted = false
			break
		}
	}
	fmt.Printf("Is s4 sorted? %v\n", isSorted)

	// ------------------- Remove duplicates -------------------
	sliceWithDup := []int{1, 2, 2, 3, 3, 3, 4}
	result := sliceWithDup[:0] // reuse underlying array
	for i, v := range sliceWithDup {
		if i == 0 || v != sliceWithDup[i-1] {
			result = append(result, v)
		}
	}
	printSlice("Remove duplicates", result)

	// ------------------- Merge two slices -------------------
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}
	merged := append(sliceA, sliceB...)
	printSlice("Merged slice", merged)

	// ------------------- Intersection -------------------
	sliceC := []int{1, 2, 3, 4}
	sliceD := []int{3, 4, 5, 6}
	intersection := []int{}
	for _, v := range sliceC {
		for _, w := range sliceD {
			if v == w {
				intersection = append(intersection, v)
				break
			}
		}
	}
	printSlice("Intersection slice", intersection)

	// ------------------- Rotate slice left by 2 -------------------
	sliceE := []int{1, 2, 3, 4, 5}
	rotateLeft := append(sliceE[2:], sliceE[:2]...)
	printSlice("Rotate left by 2", rotateLeft)

	// ------------------- Sliding window sum (size 3) -------------------
	windowSize := 3
	sliceF := []int{1, 2, 3, 4, 5}
	fmt.Print("Sliding window sums: ")
	for i := 0; i <= len(sliceF)-windowSize; i++ {
		sum := 0
		for j := i; j < i+windowSize; j++ {
			sum += sliceF[j]
		}
		fmt.Printf("%d ", sum)
	}
	fmt.Println()

	// ------------------- Sort slice -------------------
	toSort := []int{5, 2, 8, 3, 1}
	sort.Ints(toSort)
	printSlice("Sorted slice", toSort)

	// ------------------- 2D Slice -------------------
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("2D Slice (matrix):", matrix)
}

/*
	| Operation         | Example                          | Description                    |
| ----------------- | -------------------------------- | ------------------------------ |
| Declare           | `var s []int`                    | Nil slice                      |
| Create            | `s := make([]int, 5)`            | Slice of length 5, default 0   |
| Append            | `s = append(s, 10)`              | Add element at end             |
| Slice             | `sub := s[1:4]`                  | Subslice `[start:end]`         |
| Copy              | `copy(dst, src)`                 | Copies elements, returns count |
| Length            | `len(s)`                         | Number of elements             |
| Capacity          | `cap(s)`                         | Maximum without reallocation   |
| Delete            | `s = append(s[:i], s[i+1:]...)`  | Remove element at index `i`    |
| Iteration         | `for i, v := range s {}`         | Loop through slice             |
| Multi-dimensional | `matrix := [][]int{{1,2},{3,4}}` | 2D slice                       |
*/
