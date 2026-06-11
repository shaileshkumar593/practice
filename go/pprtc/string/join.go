// Go program to illustrate how to
// concatenate all the elements
// present in the slice of the string
package main

import (
	"fmt"
	"strings"
)

func main() {

	// Creating and initializing slice of string
	myslice := []string{"Welcome", "To",
		"GeeksforGeeks", "Portal"}

	// Concatenating the elements
	// present in the slice
	// Using join() function
	result := strings.Join(myslice, "-")
	fmt.Println(result)
}
