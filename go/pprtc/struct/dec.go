package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

// a struct is defined and it is not explicitly initialized with any value,
// the fields of the struct are assigned their zero values by default.
func main() {
	var emp4 Employee //zero valued struct
	fmt.Println("First Name:", emp4.firstName)
	fmt.Println("Last Name:", emp4.lastName)
	fmt.Println("Age:", emp4.age)
	fmt.Println("Salary:", emp4.salary)
}
