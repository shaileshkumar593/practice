package main

import (
	"fmt"
)

// generateSubstrings generates all substrings with unique characters
func generateSubstringss(s string) []string {
	var result []string
	runes := []rune(s) // handle Unicode properly
	n := len(runes)

	for i := 0; i < n; i++ {
		seen := make(map[rune]bool)
		substr := ""
		for j := i; j < n; j++ {
			if seen[runes[j]] {
				break // repetition found, stop extending this substring
			}
			seen[runes[j]] = true
			substr += string(runes[j])
			fmt.Println("--------------- ", substr, i, j)
			result = append(result, substr)
		}
	}

	return result
}

func main() {
	input := "abc"
	fmt.Println("Input string:", input)

	substrings := generateSubstringss(input)
	fmt.Println("All substrings without repetition:")
	for _, sub := range substrings {
		fmt.Println(sub)
	}
}
