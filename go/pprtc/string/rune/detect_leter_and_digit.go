package main

import (
	"fmt"
	"unicode"
)

func detectLettersAndDigits(s string) {
	for _, r := range s { // iterate over runes
		switch {
		case unicode.IsLetter(r):
			fmt.Printf("%c is a letter\n", r)
		case unicode.IsDigit(r):
			fmt.Printf("%c is a digit\n", r)
		default:
			fmt.Printf("%c is neither letter nor digit\n", r)
		}
	}
}

func main() {
	testStr := "GoLang123 ₹"
	detectLettersAndDigits(testStr)
}
