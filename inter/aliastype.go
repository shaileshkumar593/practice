package main

import (
	"fmt"
	"reflect"
)

type MyInts = int // Alias
type NewInt int   // New defined type
func Sum() {
	type newT int
	var s NewInt = 10
	var p newT = 20
	fmt.Println(reflect.TypeOf(s), ";;;;;;;", reflect.TypeOf(p))
}
func main() {
	var a MyInts = 10
	var b int = 20
	fmt.Println(reflect.TypeOf(a)) // int

	a = b // ✅ same type

	fmt.Println(a)

	var x NewInt = 10

	fmt.Printf("%T\n", x) // main.NewInt
	// x = b          // ❌ different types
	x = NewInt(b) // ✅ explicit conversion

	fmt.Println(x)

	Sum()
}

/*
Why main.NewInt?

Because NewInt is defined in your main package.

The fully qualified type name is:

package.Type

*/
