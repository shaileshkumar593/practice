package main

import (
	"fmt"
	"strconv"
)

func ToggleKthBit(n, k int) int {
	return n ^ (1 << (k - 1))
}
func main() {
	var n = 20
	var k = 3
	fmt.Printf("Binary of %d is: %s.\n", n, strconv.FormatInt(int64(n), 2))
	number := ToggleKthBit(n, k)
	fmt.Printf("After toggling %d rd bit of the given number is %d.\n", k, number)
}
