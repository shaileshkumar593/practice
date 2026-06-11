package main

import (
	"fmt"
	"sort"
)

func main() {

	words := []string{"bear", "atom", "world", "falcon", "cloud", "forest"}

	sort.Strings(words)
	fmt.Println(words)

	sort.Sort(sort.Reverse(sort.StringSlice(words)))
	fmt.Println(words)
}
