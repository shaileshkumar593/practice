package main

import "fmt"

func charFrequency(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}
	return freq
}

func main() {
	text := "hello world"
	freq := charFrequency(text)

	for ch, count := range freq {
		fmt.Printf("%q: %d\n", ch, count)
	}
}
