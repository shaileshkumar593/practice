package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Example 1: Valid UTF-8 string
	validStr := "Hello, 世界😊"
	fmt.Println("String:", validStr)
	if utf8.ValidString(validStr) {
		fmt.Println("utf8.ValidString: The string is valid UTF-8 ✅")
	} else {
		fmt.Println("utf8.ValidString: The string is NOT valid UTF-8 ❌")
	}

	// Convert string to []byte and validate
	validBytes := []byte(validStr)
	if utf8.Valid(validBytes) {
		fmt.Println("utf8.Valid: The byte slice is valid UTF-8 ✅")
	} else {
		fmt.Println("utf8.Valid: The byte slice is NOT valid UTF-8 ❌")
	}

	fmt.Println("---------------------------------------------------")

	// Example 2: Invalid UTF-8 byte slice
	// Manually create an invalid UTF-8 sequence
	invalidBytes := []byte{0xff, 0xfe, 0xfd}
	fmt.Println("Byte slice:", invalidBytes)

	// Convert bytes to string (unsafe)
	invalidStr := string(invalidBytes)
	fmt.Println("Converted string:", invalidStr)

	if utf8.ValidString(invalidStr) {
		fmt.Println("utf8.ValidString: The string is valid UTF-8 ✅")
	} else {
		fmt.Println("utf8.ValidString: The string is NOT valid UTF-8 ❌")
	}

	if utf8.Valid(invalidBytes) {
		fmt.Println("utf8.Valid: The byte slice is valid UTF-8 ✅")
	} else {
		fmt.Println("utf8.Valid: The byte slice is NOT valid UTF-8 ❌")
	}
}
