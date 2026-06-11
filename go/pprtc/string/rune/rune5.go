package main

import "fmt"

func main() {
	s := "₹50"

	fmt.Println("Unsafe byte slicing:")
	fmt.Println("s[:1] =", s[:1])
	fmt.Println("s[:2] =", s[:2])

	fmt.Println("\nSafe rune slicing:")
	runes := []rune(s)
	fmt.Println("runes[:1] =", string(runes[:1]))
	fmt.Println("runes[:2] =", string(runes[:2]))
}
