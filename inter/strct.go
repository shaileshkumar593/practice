package main

import "fmt"

type Employee struct {
	Name   string
	Salary int
}

func (e Employee) GetSalary() int {
	return e.Salary
}

func main() {
	emp := Employee{
		Name:   "Shailesh",
		Salary: 50000,
	}

	fmt.Println(emp.GetSalary())
}

/*

Can a struct hold methods in Go?

"Yes, we can define methods on a struct type using a receiver.
However, the method isn't physically stored in each struct instance.
The struct stores its fields, while the methods are associated with the struct's type.
We can use value receivers for read-only behavior and pointer receivers when we need to modify
the original struct or avoid copying a large struct."

*/
