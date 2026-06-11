package main

import (
	"fmt"
	"unicode"
)

func solve(word string) bool {
	charMap := make(map[rune]bool)

	for _, ch := range word {
		if unicode.IsLetter(ch) {
			if charMap[ch] {
				return false
			}
			charMap[ch] = true
		}
	}
	return true
}

func main() {
	s := "education" // all letter unique Isogram
	fmt.Println(solve(s))
}
