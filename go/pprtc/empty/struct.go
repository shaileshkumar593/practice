package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	emp1 := Employee{}
	emp2 := Employee{"John Doe", 25, 3000}

	if reflect.DeepEqual(emp1, Employee{}) {
		fmt.Println("emp1 is empty")
	} else {
		fmt.Println("emp1 is not empty")
	}

	if reflect.DeepEqual(emp2, Employee{}) {
		fmt.Println("emp2 is empty")
	} else {
		fmt.Println("emp2 is not empty")
	}
}
