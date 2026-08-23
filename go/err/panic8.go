package main

import "fmt"

func main() {
	calculate()

	fmt.Println("main continues")
}

func calculate() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("calculate started")

	panic("calculation failed")

	fmt.Println("calculate finished")
}
