package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println(ptr == nil, ptr) // true *ptr result into pannic

	/*
		In Go, pointers are a fundamental type that stores the memory address of a variable.
		When a pointer is declared but not initialized, its value is nil If you dereference a nil pointer, it will result in a panic.
		Therefore, it is crucial to check whether a pointer is nil before performing any pointer operations.

	*/

	var ptr1 *int = new(int) // declared and initialized
	fmt.Println(ptr1, *ptr1)
	fmt.Println(ptr1 == nil)

}
