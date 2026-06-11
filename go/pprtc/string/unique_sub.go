package main

import (
	"fmt"
)

// areCharactersUnique checks if all characters in a string are unique.
func areCharactersUnique(s string) bool {
	// A map to store the count of each rune (Unicode code point) in the string.
	// Go strings are UTF-8 encoded, so using rune ensures correct handling of multi-byte characters.
	charCounts := make(map[rune]int)

	for _, char := range s {
		if charCounts[char] > 0 {
			// Found a duplicate character
			return false
		}
		charCounts[char]++
	}

	// No duplicates found
	return true
}

func main() {
	string1 := "abcdefg"
	string2 := "hello world"

	fmt.Printf("'%s' has all unique characters: %v\n", string1, areCharactersUnique(string1))
	fmt.Printf("'%s' has all unique characters: %v\n", string2, areCharactersUnique(string2))
}
