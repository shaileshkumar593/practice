package main

import "fmt"

func main() {
	fmt.Println("Program started")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("database connection failed")

	fmt.Println("Program finished") // Go stops normal execution of main() and starts panic unwinding.
	// The function in which the panic occurred has already stopped executing normally.
}
