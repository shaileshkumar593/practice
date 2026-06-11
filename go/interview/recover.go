package main

import "fmt"

func main() {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}

	}()

	panic("database connection failed")

	fmt.Println("never executes")
}