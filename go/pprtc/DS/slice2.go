package main

import (
	"fmt"
	"sort"
)

func main() {
	// 21. Linear search
	s := []int{10, 20, 30, 40}
	target := 30
	found := -1
	for i, v := range s {
		if v == target {
			found = i
			break
		}
	}
	fmt.Println("21. Search:", found)

	// 22. Binary search (sorted slice)
	sort.Ints(s)
	idx := sort.SearchInts(s, 30)
	fmt.Println("22. Binary search:", idx)

	// 23. Rotate left by k
	s23 := []int{1, 2, 3, 4, 5}
	k := 2
	s23 = append(s23[k:], s23[:k]...)
	fmt.Println("23. Rotate left:", s23)

	// 24. Rotate right by k
	s24 := []int{1, 2, 3, 4, 5}
	k = 2
	s24 = append(s24[len(s24)-k:], s24[:len(s24)-k]...)
	fmt.Println("24. Rotate right:", s24)

	// 25. Merge two sorted slices
	a, b := []int{1, 3, 5}, []int{2, 4, 6}
	merged := append(a, b...)
	sort.Ints(merged)
	fmt.Println("25. Merge:", merged)

	// 26. Intersection of slices
	x := []int{1, 2, 3, 4}
	y := []int{3, 4, 5, 6}
	inter := []int{}
	set := map[int]bool{}
	for _, v := range x {
		set[v] = true
	}
	for _, v := range y {
		if set[v] {
			inter = append(inter, v)
		}
	}
	fmt.Println("26. Intersection:", inter)

	// 27. Union of slices
	union := append(x, y...)
	sort.Ints(union)
	u := []int{}
	seen := map[int]bool{}
	for _, v := range union {
		if !seen[v] {
			seen[v] = true
			u = append(u, v)
		}
	}
	fmt.Println("27. Union:", u)

	// 28. Difference (x - y)
	diff := []int{}
	check := map[int]bool{}
	for _, v := range y {
		check[v] = true
	}
	for _, v := range x {
		if !check[v] {
			diff = append(diff, v)
		}
	}
	fmt.Println("28. Difference:", diff)

	// 29. Symmetric difference
	sym := []int{}
	for _, v := range x {
		if !check[v] {
			sym = append(sym, v)
		}
	}
	for _, v := range y {
		if !set[v] {
			sym = append(sym, v)
		}
	}
	fmt.Println("29. SymDiff:", sym)

	// 30. Check if slice contains value
	contains := false
	for _, v := range x {
		if v == 3 {
			contains = true
			break
		}
	}
	fmt.Println("30. Contains 3?", contains)

	// 31. Count frequency of elements
	s31 := []int{1, 2, 2, 3, 3, 3}
	freq := map[int]int{}
	for _, v := range s31 {
		freq[v]++
	}
	fmt.Println("31. Frequency:", freq)

	// 32. Sort by frequency
	type kv struct {
		Key, Value int
	}
	var freqSlice []kv
	for k, v := range freq {
		freqSlice = append(freqSlice, kv{k, v})
	}
	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i].Value > freqSlice[j].Value
	})
	fmt.Println("32. Sort by frequency:", freqSlice)

	// 33. Remove even numbers
	s33 := []int{1, 2, 3, 4, 5, 6}
	res33 := []int{}
	for _, v := range s33 {
		if v%2 != 0 {
			res33 = append(res33, v)
		}
	}
	fmt.Println("33. Remove evens:", res33)

	// 34. Remove odd numbers
	res34 := []int{}
	for _, v := range s33 {
		if v%2 == 0 {
			res34 = append(res34, v)
		}
	}
	fmt.Println("34. Remove odds:", res34)

	// 35. Find duplicates
	s35 := []int{1, 2, 2, 3, 4, 4, 5}
	seen35 := map[int]bool{}
	dupes := []int{}
	for _, v := range s35 {
		if seen35[v] {
			dupes = append(dupes, v)
		}
		seen35[v] = true
	}
	fmt.Println("35. Duplicates:", dupes)

	// 36. Remove duplicates in place (sorted)
	s36 := []int{1, 1, 2, 2, 3, 4, 4}
	j := 0
	for i := 1; i < len(s36); i++ {
		if s36[i] != s36[j] {
			j++
			s36[j] = s36[i]
		}
	}
	s36 = s36[:j+1]
	fmt.Println("36. Unique in-place:", s36)

	// 37. Chunk slice
	s37 := []int{1, 2, 3, 4, 5, 6, 7}
	chunkSize := 3
	for i := 0; i < len(s37); i += chunkSize {
		end := i + chunkSize
		if end > len(s37) {
			end = len(s37)
		}
		fmt.Println("37. Chunk:", s37[i:end])
	}

	// 38. Flatten 2D slice
	matrix := [][]int{{1, 2}, {3, 4}, {5, 6}}
	flat := []int{}
	for _, row := range matrix {
		flat = append(flat, row...)
	}
	fmt.Println("38. Flatten:", flat)

	// 39. Compare two slices (equal?)
	a39, b39 := []int{1, 2, 3}, []int{1, 2, 3}
	equal := len(a39) == len(b39)
	if equal {
		for i := range a39 {
			if a39[i] != b39[i] {
				equal = false
				break
			}
		}
	}
	fmt.Println("39. Equal slices?", equal)

	// 40. Shuffle slice (Fisher-Yates)
	s40 := []int{1, 2, 3, 4, 5}
	for i := len(s40) - 1; i > 0; i-- {
		j := i / 2 // dummy pseudo "random"
		s40[i], s40[j] = s40[j], s40[i]
	}
	fmt.Println("40. Shuffle:", s40)
}
