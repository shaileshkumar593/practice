package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Case 1: Truly nil interface
	var x interface{} = nil
	fmt.Printf("x: (%v, %T), x == nil: %v\n", x, x, x == nil)

	// Case 2: Non-nil interface holding a nil pointer
	var p *int = nil
	var y interface{} = p
	fmt.Printf("y: (%v, %T), y == nil: %v\n", y, y, y == nil)

	// Case 3: A helper function to check for both nil conditions
	fmt.Printf("IsNil(y): %v\n", IsNil(y))
}

// IsNil checks if an interface is a nil interface or holds a nil value.
func IsNil(i interface{}) bool {
	if i == nil {
		return true // Handles truly nil interface
	}
	v := reflect.ValueOf(i)
	// Check if the underlying value kind is one that can be nil
	switch v.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface, reflect.Chan:
		return v.IsNil()
	default:
		return false
	}
}

/*
	x: (<nil>, <nil>), x == nil: true
y: (<nil>, *int), y == nil: false
IsNil(y): true
*/

/*

	var x interface{} = nil: The interface is truly nil when both its dynamic type and value are unset (both are nil). When you compare this interface to nil using x == nil, the result is true.

	var x interface{} = (*int)(nil): Here, the interface's dynamic type is set to *int, even though the underlying pointer value is nil. Because the type component is non-nil,
	 the entire interface value is considered non-nil. When you compare this interface to nil using x == nil, the result is false. */
