package main

import (
	"fmt"
	"strings" //import necessary packages fmt and strings
)

// create main function to execute the program
func main() {
	var input string //declare a string variable

	// Check if the string is empty
	if len(strings.TrimSpace(input)) == 0 {
		fmt.Println("The input string is empty or null") //if condition satisfies print the string is empty or null
	}
}
