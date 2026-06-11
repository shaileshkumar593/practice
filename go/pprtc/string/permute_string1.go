package main

import (
	"fmt"
)

// isUnique checks if all characters in s[start:end+1] are unique
func isUnique(runes []rune, start, end int) bool {
	for i := start; i < end; i++ {
		if runes[i] == runes[end] {
			return false
		}
	}
	return true
}

// generateSubstrings generates all substrings with unique characters without extra space
func generateSubstringsss(s string) []string {
	runes := []rune(s) // handle Unicode
	n := len(runes)
	var result []string

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if !isUnique(runes, i, j) {
				break
			}
			result = append(result, string(runes[i:j+1]))
		}
	}

	return result
}

func main() {
	input := "abc"
	fmt.Println("Input string:", input)

	substrings := generateSubstringsss(input)
	fmt.Println("All substrings without repetition:")
	for _, sub := range substrings {
		fmt.Println(sub)
	}
}
