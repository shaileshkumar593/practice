package main

import "fmt"

func safe() {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}

	}()

	panic("Crash")
}

func main() {
	safe()
}
