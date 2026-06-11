package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s) // convert string to runes
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i] // swap
	}

	return string(runes)
}

func main() {
	str := "Hello ₹ World 123"
	fmt.Println("Original:", str)
	fmt.Println("Reversed:", reverseString(str))
}
