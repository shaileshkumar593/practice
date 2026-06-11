package main

import (
	"fmt"
	"sort"
)

func main() {

	words := []string{"cloud", "atom", "sea", "by", "forest", "maintenance"}

	sort.Slice(words, func(i1, i2 int) bool {
		return len(words[i1]) < len(words[i2])
	})

	fmt.Println(words)

	sort.Slice(words, func(i1, i2 int) bool {
		return len(words[i1]) > len(words[i2])
	})

	fmt.Println(words)
}
