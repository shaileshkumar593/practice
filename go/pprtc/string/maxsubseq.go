package main

import "fmt"

// longestUniqueSubstring prints iteration windows and returns the longest substring
func longestUniqueSubstring(s string) string {
	runes := []rune(s)
	n := len(runes)
	start, maxLen := 0, 0
	begin := 0

	for end := 0; end < n; end++ {
		// Check for duplicates in current window
		for i := start; i < end; i++ {
			if runes[i] == runes[end] {
				start = i + 1
				break
			}
		}

		// Update max substring info
		if end-start+1 > maxLen {
			maxLen = end - start + 1
			begin = start
		}

		// Print current window
		fmt.Printf("Window [%d:%d] -> \"%s\"\n", start, end, string(runes[start:end+1]))
	}

	fmt.Printf("\nLongest substring window [%d:%d] -> \"%s\"\n", begin, begin+maxLen-1, string(runes[begin:begin+maxLen]))
	return string(runes[begin : begin+maxLen])
}

func main() {
	s := "aababcdc"
	fmt.Println("Input string:", s)
	longest := longestUniqueSubstring(s)
	fmt.Println("Longest substring without repeating characters:", longest)
}
