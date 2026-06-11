package main

import (
	"fmt"
)

type Person struct {
	string
	int
}

// Even though anonymous fields do not have an explicit name,
// by default the name of an anonymous field is the name of its type
func main() {
	p1 := Person{
		string: "naveen",
		int:    50,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
