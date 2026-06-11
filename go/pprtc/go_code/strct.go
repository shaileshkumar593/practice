package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	x := person{
		age: 30, firstName: "Kevin", lastName: "Jhon",
	}

	fmt.Println(x)
	fmt.Println(x.firstName)
	fmt.Println(x.lastName)
}
