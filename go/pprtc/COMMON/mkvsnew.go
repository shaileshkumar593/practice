package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := new(Person)

	a := make([]int, 2, 2)

	fmt.Println(a)
	fmt.Println(p)
}
