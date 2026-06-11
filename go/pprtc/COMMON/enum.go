package main

import (
	"fmt"
)

// Enumerated Constant

const (
	_ = iota + 96
	a
	b
	c
	d
	e
)

func main() {

	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(string(c))
	fmt.Println(string(d))
	fmt.Println(string(e))

}
