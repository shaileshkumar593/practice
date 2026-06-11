package main

import "fmt"

/*
	A nil slice does not point to any valid underlying array, and both its length(len) and capacity(cap) are 0.
	 However, a nil slice and an empty slice(created with make([]int, 0) or []int{}) are different.
	 A nil slice does not occupy memory until space is allocated for it, while an empty slice, although having a length of 0,
	 already has a pointer pointing to an underlying array with a length of 0.

*/

func main() {
	var s []int                 //slice
	var p [6]int                // array
	t := [...]int{}             //array
	fmt.Println(s == nil, p, t) // true

	// nil Map
	/*
		A map is used to store a collection of key-value pairs, where keys are unique. When a map is declared but not initialized,
		its value is nil. This implies that no memory space has been allocated, and it cannot be used directly.

	*/
	var myMap map[string]int
	fmt.Println(myMap == nil)

}
