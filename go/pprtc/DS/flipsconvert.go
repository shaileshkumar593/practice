package main

import (
	"fmt"
	"strconv"
)

func FindBits(x, y int) int {
	n := x ^ y
	count := 0
	for ; n != 0; count++ {
		n = n & (n - 1)
	}
	return count
}
func main() {
	x := 65
	y := 80
	fmt.Printf("Binary of %d is: %s.\n", x, strconv.FormatInt(int64(x), 2))
	fmt.Printf("Binary of %d is: %s.\n", y, strconv.FormatInt(int64(y), 2))
	fmt.Printf("The number of bits flipped is %d\n", FindBits(x, y))
}
