package main

import "fmt"

//Calling a nil function will result in a panic.
func main() {
	var fn func(int) int
	fmt.Println(fn == nil) // true
}
