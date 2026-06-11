package main

import "fmt"

func longestUniqueSubstring(s string) string {
	runes := []rune(s)
	n := len(runes)
	start, maxLen := 0, 0
	var pos [256]int // tracks last seen index, 0 default
	var begin int

	for end := 0; end < n; end++ {
		c := runes[end]
		fmt.Println("c    ", c)
		if last := pos[c]; last > begin {
			fmt.Println("last   ", pos[c])
			begin = last
		}
		pos[c] = end + 1
		if end-begin+1 > maxLen {
			maxLen = end - begin + 1
			start = begin
		}
	}
	return string(runes[start : start+maxLen])
}

func main() {
	s := "abcabcbb"
	fmt.Println("Longest substring without repeating characters:", longestUniqueSubstring(s))
}
