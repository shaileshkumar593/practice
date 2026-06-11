package main

import (
	"fmt"
)

// hasDuplicate checks if a rune exists in the slice
func hasDuplicate(slice []rune, char rune) bool {
	for _, r := range slice {
		if r == char {
			return true
		}
	}
	return false
}

// generateSubstrings generates all substrings with unique characters without using map
func generateSubstrings(s string) []string {
	var result []string
	runes := []rune(s) // handle Unicode safely
	n := len(runes)

	for i := 0; i < n; i++ {
		var substrRunes []rune
		for j := i; j < n; j++ {
			if hasDuplicate(substrRunes, runes[j]) {
				break // repetition found, stop extending this substring
			}
			substrRunes = append(substrRunes, runes[j])
			result = append(result, string(substrRunes))
		}
	}

	return result
}

func main() {
	input := "abc"
	fmt.Println("Input string:", input)

	substrings := generateSubstrings(input)
	fmt.Println("All substrings without repetition:")
	for _, sub := range substrings {
		fmt.Println(sub)
	}
}
