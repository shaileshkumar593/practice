package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "Hello World! 123 ₹"

	var vowels, consonants int

	for _, r := range s {
		if unicode.IsLetter(r) { // Check if it's a letter
			switch unicode.ToLower(r) {
			case 'a', 'e', 'i', 'o', 'u':
				vowels++
			default:
				consonants++
			}
		}
	}

	fmt.Println("Input string:", s)
	fmt.Println("Vowels count:", vowels)
	fmt.Println("Consonants count:", consonants)
}
