package main

import (
	"fmt"
	"unicode"
)

func main() {
	r := '1'
	if unicode.IsLetter(r) {
		fmt.Println("The rune is a letter")
	} else if unicode.IsDigit(r) {
		fmt.Println("The rune is a digit")
	} else {
		fmt.Println("The rune is neither a letter nor a digit")
	}
}
