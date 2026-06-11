package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("----- Slice Toolbox in Go -----")

	// 1. Basic slice creation
	s := []int{3, 1, 4, 1, 5}
	fmt.Println("Original slice:", s)

	// 2. len and cap
	fmt.Println("Length:", len(s), "Capacity:", cap(s))

	// 3. Append elements
	s = append(s, 9, 2)
	fmt.Println("After append:", s)

	// 4. Delete element at index 2
	i := 2
	s = append(s[:i], s[i+1:]...)
	fmt.Println("After delete index 2:", s)

	// 5. Insert element at index 2
	val := 8
	s = append(s[:i], append([]int{val}, s[i:]...)...)
	fmt.Println("After insert 8 at index 2:", s)

	// 6. Copy slices
	dst := make([]int, len(s))
	copy(dst, s)
	fmt.Println("Copied slice:", dst)

	// 7. Slice sub-slice
	sub := s[1:4]
	fmt.Println("Sub-slice [1:4]:", sub)

	// 8. Reset slice but keep capacity
	s = s[:0]
	fmt.Println("Reset slice:", s, "Capacity remains:", cap(s))

	// 9. Sort slices
	s2 := []int{5, 2, 9, 1}
	sort.Ints(s2)
	fmt.Println("Sorted int slice:", s2)

	// 10. Sort strings
	strSlice := []string{"banana", "apple", "cherry"}
	sort.Strings(strSlice)
	fmt.Println("Sorted string slice:", strSlice)

	// 11. Custom sort
	sort.Slice(strSlice, func(i, j int) bool {
		return len(strSlice[i]) < len(strSlice[j])
	})
	fmt.Println("Custom sort by length:", strSlice)

	// 12. Strings -> slices
	word := "hello"
	byteSlice := []byte(word)
	runeSlice := []rune(word)
	fmt.Println("Byte slice:", byteSlice)
	fmt.Println("Rune slice:", runeSlice)

	// 13. Split and join
	csv := "a,b,c"
	parts := strings.Split(csv, ",")
	fmt.Println("Split:", parts)
	joined := strings.Join(parts, "-")
	fmt.Println("Join:", joined)

	// 14. Bytes package for []byte slices
	data := []byte("GoLang")
	splitBytes := bytes.Split(data, []byte("a"))
	fmt.Println("Bytes split:", splitBytes)
	joinedBytes := bytes.Join(splitBytes, []byte("_"))
	fmt.Println("Bytes join:", joinedBytes)
}

/*

	| Operation           | Function / Idiom                                       | Description                      |
	| ------------------- | ------------------------------------------------------ | -------------------------------- |
	| Create slice        | `[]T{}` or `make([]T, len, cap)`                       | Initialize slice                 |
	| Length              | `len(slice)`                                           | Number of elements               |
	| Capacity            | `cap(slice)`                                           | Maximum elements before resize   |
	| Append              | `append(slice, elems...)`                              | Add elements at end              |
	| Delete              | `slice[:i] + slice[i+1:]`                              | Remove element at index i        |
	| Insert              | `append(slice[:i], append([]T{val}, slice[i:]...)...)` | Insert element at index i        |
	| Copy                | `copy(dst, src)`                                       | Copy slice elements              |
	| Sub-slice           | `slice[i:j]`                                           | Slice from i to j-1              |
	| Reset               | `slice[:0]`                                            | Reset length to 0, keep capacity |
	| Sort ints           | `sort.Ints(slice)`                                     | Sort int slice ascending         |
	| Sort strings        | `sort.Strings(slice)`                                  | Sort string slice ascending      |
	| Custom sort         | `sort.Slice(slice, func)`                              | Sort using custom comparator     |
	| String → byte slice | `[]byte(s)`                                            | Convert string to byte slice     |
	| String → rune slice | `[]rune(s)`                                            | Convert string to rune slice     |
	| Split               | `strings.Split(s, sep)`                                | Split string into []string       |
	| Join                | `strings.Join(slice, sep)`                             | Join []string into string        |
	| Bytes Split         | `bytes.Split(b, sep)`                                  | Split []byte by separator        |
	| Bytes Join          | `bytes.Join(slice, sep)`                               | Join [][]byte into []byte        |
*/

/*
	Standard Library Functions That Work With Slices
1. sort package
		sort.Ints(slice []int) → sorts int slice ascending

		sort.Strings(slice []string) → sorts string slice

		sort.Float64s(slice []float64) → sorts float64 slice

		sort.Slice(slice, func(i,j int) bool) → custom sort

*/
