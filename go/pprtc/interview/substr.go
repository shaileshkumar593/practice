package main

import "fmt"

func canForm(first, second string) bool {
	freq := make(map[rune]int)

	// Count frequencies in first string
	for _, ch := range first {
		freq[ch]++
	}

	// Check if we can form second
	for _, ch := range second {
		freq[ch]--
		if freq[ch] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canForm("aabbc", "abc"))    // true
	fmt.Println(canForm("aabbc", "aabc"))   // true
	fmt.Println(canForm("aabbc", "abbc"))   // false
	fmt.Println(canForm("hello", "helloo")) // false
}
