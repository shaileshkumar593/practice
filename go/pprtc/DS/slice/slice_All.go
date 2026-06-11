package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera"}
	for i, v := range slices.All(names) {
		fmt.Println(i, ":", v)
	}
}

/*

	All returns an iterator over index-value pairs in the slice in the usual order.
	Output:

		0 : Alice
		1 : Bob
		2 : Vera
*/
