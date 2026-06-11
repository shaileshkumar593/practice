package main

import (
	"fmt"
)

func main() {
	fmt.Println("==== ASCII Example ====")

	s1 := "hello"

	// 🔹 1. Index-based loop (byte)
	fmt.Println("\n-- Index Loop (byte) --")
	for i := 0; i < len(s1); i++ {
		b := s1[i]
		fmt.Printf("index=%d, value=%d, char=%c, type=%T\n", i, b, b, b)
	}

	// 🔹 2. Range loop (rune)
	fmt.Println("\n-- Range Loop (rune) --")
	for i, ch := range s1 {
		fmt.Printf("index=%d, value=%d, char=%c, type=%T\n", i, ch, ch, ch)
	}

	// ----------------------------------------------------

	fmt.Println("\n==== Unicode Example ====")

	s2 := "héllo"

	// 🔹 3. Index loop (byte - WRONG for Unicode)
	fmt.Println("\n-- Index Loop (byte) --")
	for i := 0; i < len(s2); i++ {
		b := s2[i]
		fmt.Printf("index=%d, value=%d, char=%c, type=%T\n", i, b, b, b)
	}

	// 🔹 4. Range loop (rune - CORRECT)
	fmt.Println("\n-- Range Loop (rune) --")
	for i, ch := range s2 {
		fmt.Printf("index=%d, value=%d, char=%c, type=%T\n", i, ch, ch, ch)
	}

	// ----------------------------------------------------

	// 🔹 5. Convert string to rune slice (safe indexing)
	fmt.Println("\n-- Rune Slice Indexing --")
	runes := []rune(s2)

	for i := 0; i < len(runes); i++ {
		fmt.Printf("index=%d, value=%d, char=%c, type=%T\n", i, runes[i], runes[i], runes[i])
	}
}