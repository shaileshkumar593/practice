package main

import "fmt"

// Check if character c exists in substring s[0:end]
func hasDuplicate(s []rune, end int) bool {
	for i := 0; i < end; i++ {
		if s[i] == s[end] {
			return true
		}
	}
	return false
}

// Recursive function to find longest substring without repeating chars
func longestUniqueSubstringRec(s []rune, start int) string {
	if start >= len(s) {
		return ""
	}

	maxStr := ""
	for end := start; end < len(s); end++ {
		if hasDuplicate(s[start:end+1], end-start) {
			break
		}
		current := string(s[start : end+1])
		if len(current) > len(maxStr) {
			maxStr = current
		}
	}

	// Recurse for the substring starting from next character
	rest := longestUniqueSubstringRec(s, start+1)
	if len(rest) > len(maxStr) {
		return rest
	}
	return maxStr
}

func main() {
	input := "abcabcbb"
	runes := []rune(input)
	result := longestUniqueSubstringRec(runes, 0)
	fmt.Println("Longest substring without repeating characters:", result)
}
