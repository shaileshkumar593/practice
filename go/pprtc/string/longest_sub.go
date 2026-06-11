package main

import "fmt"

// longestUniqueSubstringOptimized finds longest substring without repeating characters in O(n)
func longestUniqueSubstringOptimized(s string) string {
	runes := []rune(s)
	n := len(runes)
	start, maxLen, maxStart := 0, 0, 0
	lastSeen := make(map[rune]int) // stores last index of each rune

	for end := 0; end < n; end++ {
		if idx, found := lastSeen[runes[end]]; found && idx >= start {
			start = idx + 1 // move start to exclude duplicate
		}

		lastSeen[runes[end]] = end // update last seen index

		// Update max substring info
		if end-start+1 > maxLen {
			maxLen = end - start + 1
			maxStart = start
		}

		// Print current window
		fmt.Printf("Window [%d:%d] -> \"%s\"\n", start, end, string(runes[start:end+1]))
	}

	fmt.Printf("\nLongest substring window [%d:%d] -> \"%s\"\n", maxStart, maxStart+maxLen-1, string(runes[maxStart:maxStart+maxLen]))
	return string(runes[maxStart : maxStart+maxLen])
}

func main() {
	s := "abcacbcbb"
	fmt.Println("Input string:", s)
	longest := longestUniqueSubstringOptimized(s)
	fmt.Println("Longest substring without repeating characters:", longest)
}
