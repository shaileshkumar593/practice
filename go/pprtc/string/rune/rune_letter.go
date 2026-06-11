package main

import "fmt"

func main() {
	r := 'A'
	if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' {
		fmt.Println("The rune is a letter")
	} else {
		fmt.Println("The rune is not a letter")
	}
}
