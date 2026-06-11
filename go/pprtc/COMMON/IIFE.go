package main

import "fmt"

// Immediately Invoked Function Expression (IIFE) in Go.
func main() {
	result := func(a, b int) int {
		return a + b
	}(10, 20) // immediately invoked

	fmt.Println("10 + 20 =", result)
}

/*

	Why Use IIFE in Go?
		Encapsulation: You can execute code without polluting the outer scope with temporary variables.

		Immediate computation: Useful when you need a value computed once at the point of declaration.

		Closures: Can capture surrounding variables for temporary use.

*/
