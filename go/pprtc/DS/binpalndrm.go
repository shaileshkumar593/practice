package main

import (
	"fmt"
	"strconv"
)

func IsPalindrome(n int) bool {
	rev := 0
	k := n
	for k != 0 {
		rev = (rev << 1) | (k & 1)
		k = k >> 1
	}
	return n == rev
}
func main() {
	n := 3
	fmt.Printf("Binary representation of %d is: %s.\n", n,
		strconv.FormatInt(int64(n), 2))
	if IsPalindrome(n) == true {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a Palindrome")
	}
}
