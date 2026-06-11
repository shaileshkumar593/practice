package main

import (
	"fmt"
	"reflect"
)

// reflect.DeepEqual use for value comparision
func main() {

	a1 := []int{1, 2, 3}
	b1 := []int{1, 2, 3}

	fmt.Println(a1 == nil)

	equal1 := reflect.DeepEqual(a1, b1)
	fmt.Println(equal1)
	a := map[string]int{"one": 1, "two": 2}
	b := map[string]int{"one": 1, "two": 2}

	equal := reflect.DeepEqual(a, b)
	fmt.Println(equal) // Output: true

	d1 := make([]int, 0)
	d2 := make([]int, 0)
	fmt.Println(d1 == nil, d2 == nil)

	//fmt.Println(d1 == d2) invalid operation: d1 == d2 (slice can only be compared to nil)
}
