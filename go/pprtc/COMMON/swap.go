package main

import (
	"fmt"
)

func swap_values(a *int, b *int) {

	var tmp = *a

	*a = *b

	*b = tmp

}

func main() {

	var a = 4

	var b = 10

	fmt.Printf("Before swap: a = %d, b = %d \n", a, b)

	// swap the values
	swap_values(&a, &b)

	fmt.Printf("After swap: a = %d, b = %d \n", a, b)
	var u, v = 20, 30
	v, u = u, v
	fmt.Printf("After swap: u = %d, v = %d \n", u, v)

	var r, s = "ram", "shyam"
	fmt.Printf("Before swap: r = %s, s = %s \n", r, s)
	s, r = r, s
	fmt.Printf("After swap: r = %s, s =%s", r, s)

}
