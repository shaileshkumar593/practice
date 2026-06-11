package main

import (
	"fmt"
)

func sortStringChars(s string) string {
	// Convert string to slice of runes (handles Unicode)
	chars := []rune(s)

	// Insertion sort
	for i := 1; i < len(chars); i++ {
		key := chars[i]
		j := i - 1
		for j >= 0 && chars[j] > key {
			chars[j+1] = chars[j]
			j--
		}
		chars[j+1] = key
	}

	return string(chars)
}

func main() {
	str := "golang"
	sortedStr := sortStringChars(str)
	fmt.Println("Sorted characters:", sortedStr)
}
