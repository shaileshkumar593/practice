package main

import (
	"fmt"
	"strconv"
)

func IntegerToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}
func main() {
	n := 1
	fmt.Printf("Binary Representation of %d is %s.\n", n, IntegerToBinary(n))
	n = 5
	fmt.Printf("Binary Representation of %d is %s.\n", n, IntegerToBinary(n))
	n = 20
	fmt.Printf("Binary Representation of %d is %s.\n", n, IntegerToBinary(n))
	n = 31
	fmt.Printf("Binary Representation of %d is %s.\n", n, IntegerToBinary(n))
}
