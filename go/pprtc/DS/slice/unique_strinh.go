package main

import (
	"fmt"
)

func removeRepeatingChars(s string) string {
	runes := []rune(s) // convert string to rune slice to handle Unicode
	n := len(runes)

	if n < 2 {
		return s
	}

	// i is the index for the next unique character
	i := 0

	for j := 0; j < n; j++ {
		duplicate := false
		// check if runes[j] already exists in runes[0..i-1]
		for k := 0; k < i; k++ {
			if runes[k] == runes[j] {
				duplicate = true
				break
			}
		}
		if !duplicate {
			runes[i] = runes[j]
			i++
		}
	}

	// convert the unique portion back to string
	return string(runes[:i])
}

func main() {
	s := "programming" ///
	fmt.Println("Original:", s)
	result := removeRepeatingChars(s)
	fmt.Println("After removing repeating chars:", result)

	s2 := "education"
	fmt.Println("\nOriginal:", s2)
	result2 := removeRepeatingChars(s2)
	fmt.Println("After removing repeating chars:", result2)
}

/*
	Original: programming
	After removing repeating chars: progamin

	Original: education
	After removing repeating chars: education
*/
