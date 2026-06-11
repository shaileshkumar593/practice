package main

import "fmt"

// printSlidingWindows prints all sliding windows and the longest unique substring
func printSlidingWindows(s string) string {
	runes := []rune(s)
	n := len(runes)
	start, maxLen := 0, 0
	var longest string

	for end := 0; end < n; end++ {
		// Check for duplicates in current window
		for i := start; i < end; i++ {
			if runes[i] == runes[end] {
				start = i + 1
				break
			}
		}

		// Print current window
		fmt.Printf("Window [%d:%d] -> \"%s\"\n", start, end, string(runes[start:end+1]))

		// Update longest substring if needed
		if end-start+1 > maxLen {
			maxLen = end - start + 1
			longest = string(runes[start : end+1])
		}
	}

	fmt.Printf("\nLongest substring -> \"%s\"\n", longest)
	return longest
}

func main() {
	s := "abcac"
	fmt.Println("Input string:", s)
	longest := printSlidingWindows(s)
	fmt.Println("Longest substring without repeating characters:", longest)
}
