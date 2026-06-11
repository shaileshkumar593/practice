package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	} else {
		return a
	}
}

func main() {
	a := 36
	b := 60

	fmt.Println(gcd(a, b))

}
