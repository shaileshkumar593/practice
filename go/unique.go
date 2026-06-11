package main

import "fmt"

func main() {

	str := "abcab"

	var result byte = 0

	for i := 0; i < len(str); i++ {
		result ^= str[i]
	}

	fmt.Printf("Unique char: %c\n", result)
}
