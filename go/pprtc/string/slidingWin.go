package main

import "fmt"

// generateSubstringsOptimized generates all substrings with unique characters
func generateSubstringsOptimized(s string) []string {
	runes := []rune(s)
	n := len(runes)
	var result []string

	for i := 0; i < n; i++ {
		// Use a bitmask or slice to check duplicates (0 extra space if small charset)
		for j := i; j < n; j++ {
			unique := true
			// Check if runes[j] is repeated in runes[i:j]
			for k := i; k < j; k++ {
				if runes[k] == runes[j] {
					unique = false
					break
				}
			}
			if !unique {
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

	substrings := generateSubstringsOptimized(input)
	fmt.Println("All substrings without repetition:")
	for _, sub := range substrings {
		fmt.Println(sub)
	}
}
