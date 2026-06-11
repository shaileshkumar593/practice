package main

import "fmt"

func mutate(s string) string {
	b := []byte(s) // convert to byte slice
	b[0] = 'a'     // mutate first character
	return string(b)
}

func main() {
	h := "hello"
	fmt.Println(mutate(h)) // prints: "aello"
}
