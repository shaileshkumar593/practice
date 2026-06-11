// returning anonymous function

package main

import "fmt"

func retmultival(p, q int) (int, int, int) {
	return p - q, p + q, p * q
}

func main() {
	m1, m2, m3 := retmultival(10, 8)
	fmt.Println("Return value of function", m1, m2, m3)
}
