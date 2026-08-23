package main

import (
	"fmt"
)

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Steve"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

/*
	When a function encounters a panic, its execution is stopped, any deferred functions are executed and then
	the control returns to its caller. This process continues until all the functions of the current goroutine
	have returned at which point the program prints the panic message, followed by the stack trace and then terminates.*/
