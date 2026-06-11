package main

import (
	"fmt"
	"unicode"
)

func main() {
	r := 'A'
	if unicode.IsLetter(r) {
		fmt.Println("The rune is a letter")
	} else {
		fmt.Println("The rune is not a letter")
	}
}
