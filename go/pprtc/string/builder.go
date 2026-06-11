package main

import (
	"fmt"
	"strings"
)

func main() {
	var sb strings.Builder

	sb.WriteString("Hello")
	sb.WriteString(" ")
	sb.WriteString("World")
	sb.WriteRune('!')

	finalString := sb.String()
	fmt.Println(finalString) // Output: Hello World!

	sb.Reset()  // Clear the builder for reuse
	sb.Grow(20) // Pre-allocate capacity
	sb.WriteString("Go is awesome.")
	fmt.Println(sb.String()) // Output: Go is awesome.
}
