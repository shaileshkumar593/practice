package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define a boolean flag named "is_admin" with a default value of false
	isAdmin := flag.Bool("is_admin", false, "specifies whether the user is an admin or not")

	// Parse the command-line arguments
	flag.Parse()

	// Print the value of the flag
	fmt.Println("Is Admin:", *isAdmin)
}

/*

	The syntax of the flag.Bool() function is as follows ?

		flag.Bool(name string, default bool, usage string) *bool

	name ? A string that specifies the name of the flag.

	default ? A boolean value that specifies the default value of the flag.

	usage ? A string that specifies the usage message of the flag.

*/
