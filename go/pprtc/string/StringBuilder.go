/*Using strings.Builder: You can also create a string by concatenating the strings using strings.
Builder with WriteString() method. It is defined under strings package.
It uses less memory while concatenating strings and is a better way of concatenation.
The function WriteString appends the contents of the string to a buffer and allows us to concatenate strings
in a faster way. It prevents the generation of the unnecessary string object, means
it doesn’t generate a new string like in + operator from two or more strings.*/

// Go program to illustrate how to concatenate strings
// Using strings.Builder with WriteString() function

package main

import (
	"fmt"
	"strings"
)

func main() {

	// Creating and initializing strings
	// Using strings.Builder with
	// WriteString() function
	var str strings.Builder

	str.WriteString("Welcome")

	fmt.Println("String: ", str.String())

	str.WriteString(" to")
	str.WriteString(" GeeksforGeeks!")

	fmt.Println("String: ", str.String())

}
