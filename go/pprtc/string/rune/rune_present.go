package main

/*
	Checking for a specified rune in a Golang string is relatively simple. We can use the strings.IndexRune function to check if a specific rune is present in a given string.

Here is the syntax of the strings.IndexRune function ?

func IndexRune(s string, r rune) int
The IndexRune function takes two arguments: s and r. s is the string in which we want to search for the specified rune, and r is the rune we want to find.

The function returns the index of the first occurrence of the specified rune in the given string. If the specified rune is not found in the string, the function returns -1.

*/

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, World!"
	r := 'l'

	index := strings.IndexRune(s, r)

	if index == -1 {
		fmt.Printf("Rune '%c' not found in string '%s'\n", r, s)
	} else {
		fmt.Printf("Rune '%c' found at index %d in string '%s'\n", r, index, s)
	}
}
