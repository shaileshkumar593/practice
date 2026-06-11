package main

import (
	"fmt"
	"reflect"
)

/*

 use the reflect.DeepEqual() function that compares two values recursively,
 which means it traverses and checks the equality of the corresponding data values at each level.
 However, this solution is much slower and less safe than comparing values in a loop.
 A general rule of thumb when using the reflect package is: it should be used with care and avoided unless strictly necessary.

*/

func main() {

	a := []string{"strawberry", "raspberry"}
	b := []string{"strawberry", "raspberry"}
	fmt.Printf("slices equal: %t\n", reflect.DeepEqual(a, b))

	b = append(b, "test")
	fmt.Printf("slices equal: %t\n", reflect.DeepEqual(a, b))

}
