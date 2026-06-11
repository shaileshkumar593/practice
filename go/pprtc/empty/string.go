package main

import (
	"fmt"
)

// create main function to execute the program
func main() {
	var value string //declare a string variable

	if value == "" {
		fmt.Println("The input string is empty")
	}
	if value == "" || value == "null" {
		fmt.Println("The input string is null") //if it satisfies the if statement print the string is null
	}
}
