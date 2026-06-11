package main

import "fmt"

func mutate(s string) string {
	r := []rune(s)
	r[0] = 'आ' // replace first character safely
	return string(r)
}

func main() {
	h := "नमस्ते"
	fmt.Println(mutate(h)) // prints: "आमस्ते"
}

/*

	If you want to mutate or change characters, you need to:

	Convert the string to a mutable byte slice ([]byte),
	or a rune slice ([]rune if dealing with Unicode characters).

	Modify the slice.

	Convert it back to a string.

*/
