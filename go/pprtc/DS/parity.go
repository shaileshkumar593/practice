package main

import (
	"fmt"
	"strconv"
)

func FindParity(n int) int {
	parity := 0
	p := n
	q := n
	fmt.Println("....", p<<1) // multiplication by 2
	fmt.Println("----", q>>1) // division by 2
	for n != 0 {
		fmt.Println(n & 1)
		if n&1 != 0 {
			parity += 1
			fmt.Println(parity)
		}
		n = n >> 1 // right shift
	}
	return parity
}
func main() {
	n := 20
	fmt.Printf("Binary of %d is: %s.\n", n, strconv.FormatInt(int64(n), 2))
	fmt.Println("-------->", n)
	if FindParity(n)%2 != 0 {
		fmt.Printf("Parity of the %d is Odd.\n", n)
	} else {
		fmt.Printf("Parity of the %d is Even.\n", n)
	}
}
