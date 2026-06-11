package main

import (
	"fmt"
)

type Vehicle struct {
}

func main() {

	vehicle := Vehicle{}

	switch {

	case vehicle == Vehicle{}:

		fmt.Println("Structure is empty")

	default:

		fmt.Println("Structure is not empty")
	}
}
