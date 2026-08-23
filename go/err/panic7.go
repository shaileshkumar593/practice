package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	calculate()

	fmt.Println("main continues")
}

func calculate() {
	fmt.Println("calculate started")

	panic("calculation failed")

	fmt.Println("calculate finished")
}

/*

panic abandons normal execution of the panicking function and starts stack unwinding.
recover can stop the panic during that unwinding, but it does not resume execution at
the point where the panic occurred.*/

// main() does NOT resume at the line where calculate() was called.

/*
	recover() does not mean "continue after the panic." It means "stop the panic during unwinding."
	Where the recovery happens determines which function can return normally to its caller.
*/
