package main

import (
	"fmt"
	"strings"
)

func wordFrequency(s string) map[string]int {
	freq := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		freq[word]++
	}

	return freq
}

func main() {
	text := "go is fun and go is fast"
	freq := wordFrequency(text)

	for word, count := range freq {
		fmt.Printf("%s: %d\n", word, count)
	}
}
