package main

import (
	"fmt"
	"sort"
)

func main() {

	words := []string{"bear", "atom", "world", "falcon", "cloud", "forest"}

	sorted := sort.StringsAreSorted(words)
	fmt.Println(sorted)

	sort.Sort(sort.StringSlice(words))

	sorted = sort.StringsAreSorted(words)
	fmt.Println(sorted)

	// --------------------------------------------

	vals := []int{5, 2, 6, 3, 1, 7, 8, 4, 0}

	sorted = sort.IntsAreSorted(vals)
	fmt.Println(sorted)

	sort.Ints(vals)

	sorted = sort.IntsAreSorted(vals)
	fmt.Println(sorted)
}
