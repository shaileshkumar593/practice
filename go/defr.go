package main

import (
	"fmt"
)

// function taking two arguments and returning sum

/*

	Why Go forces “function only”
	Because defer internally stores a deferred call frame, like:

	function + arguments + execution context
	So Go needs something executable like:

	✔ func call
	✔ method call
	✔ anonymous function

	NOT:

	raw expressions

	statements

	assignments

*/

func add_values(a int, b int) (i int) {
	//defer i+1 /// error defer expect function call

	defer func() {
		i = i + 1
	}()
	var c = a + b
	return c
}

// main function

func main() {

	// defining variables
	var a, b, c int

	// taking first input from user
	fmt.Scan(&a)

	// taking second input from user
	fmt.Scan(&b)

	// calling function add_values with two arguments
	c = add_values(a, b)

	// printing the sum
	fmt.Printf("Sum: %d ", c)

}
