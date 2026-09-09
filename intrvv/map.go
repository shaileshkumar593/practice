package main

import (
	"fmt"
)

func main() {

	var m map[string]int

	m["k"] = 90 // panic

	fmt.Println(m["k"])

}
