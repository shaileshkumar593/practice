package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1. Create a slice
	s1 := []int{1, 2, 3}
	fmt.Println("1. Slice:", s1)

	// 2. Create a slice using make
	s2 := make([]int, 5) // length=5
	fmt.Println("2. Slice with make:", s2)

	// 3. Append elements
	s3 := []int{1, 2}
	s3 = append(s3, 3, 4)
	fmt.Println("3. Append:", s3)

	// 4. Concatenate slices
	s4 := append([]int{1, 2}, []int{3, 4}...)
	fmt.Println("4. Concatenate:", s4)

	// 5. Iterate with index
	for i := range s4 {
		fmt.Println("5. Iterate index:", i, s4[i])
	}

	// 6. Iterate with value
	for _, v := range s4 {
		fmt.Println("6. Iterate value:", v)
	}

	// 7. Copy slice
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("7. Copy:", dst)

	// 8. Slice operator [:]
	s5 := []int{1, 2, 3, 4, 5}
	fmt.Println("8. Sub-slice:", s5[1:4])

	// 9. Delete element at index 2
	s6 := []int{10, 20, 30, 40}
	s6 = append(s6[:2], s6[3:]...)
	fmt.Println("9. Delete index 2:", s6)

	// 10. Insert at index
	s7 := []int{1, 2, 4, 5}
	index := 2
	s7 = append(s7[:index], append([]int{3}, s7[index:]...)...)
	fmt.Println("10. Insert:", s7)

	// 11. Reverse slice
	s8 := []int{1, 2, 3, 4}
	for i, j := 0, len(s8)-1; i < j; i, j = i+1, j-1 {
		s8[i], s8[j] = s8[j], s8[i]
	}
	fmt.Println("11. Reverse:", s8)

	// 12. Check if slice is empty
	var s9 []int
	fmt.Println("12. Empty?", len(s9) == 0)

	// 13. Unique elements (with map)
	s10 := []int{1, 2, 2, 3, 4, 4}
	unique := []int{}
	seen := map[int]bool{}
	for _, v := range s10 {
		if !seen[v] {
			unique = append(unique, v)
			seen[v] = true
		}
	}
	fmt.Println("13. Unique:", unique)

	// 14. Find max
	s11 := []int{10, 3, 5, 7}
	max := s11[0]
	for _, v := range s11 {
		if v > max {
			max = v
		}
	}
	fmt.Println("14. Max:", max)

	// 15. Find min
	min := s11[0]
	for _, v := range s11 {
		if v < min {
			min = v
		}
	}
	fmt.Println("15. Min:", min)

	// 16. Sum of slice
	sum := 0
	for _, v := range s11 {
		sum += v
	}
	fmt.Println("16. Sum:", sum)

	// 17. Average
	avg := float64(sum) / float64(len(s11))
	fmt.Println("17. Average:", avg)

	// 18. Sort ascending
	s12 := []int{5, 1, 4, 2}
	sort.Ints(s12)
	fmt.Println("18. Sort Asc:", s12)

	// 19. Sort descending
	sort.Sort(sort.Reverse(sort.IntSlice(s12)))
	fmt.Println("19. Sort Desc:", s12)

	// 20. Check if sorted
	isSorted := sort.IntsAreSorted(s12)
	fmt.Println("20. Sorted?", isSorted)
}
