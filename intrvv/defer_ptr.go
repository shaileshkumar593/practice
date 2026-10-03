package main

import (
	"fmt"
)

func main() {
	x := 10

	defer fmt.Println(x) // passed value is evaluated
	defer func(p *int) {
		fmt.Println(*p)
	}(&x)

	defer func(x int) { // passes valued is evaluated
		fmt.Println(x)
	}(x)

	defer func() { // closure count
		fmt.Println(x)
	}()

	x = 20

	s := [6]int{25, 24, 1, 45, 36, 44}

	fmt.Println(len(s), cap(s))
	p := s[1:4]
	fmt.Println(len(p), cap(p))
	p = append(p, 124)
	fmt.Println(p)
	fmt.Println(s)
}

// 20, 10
//The pointer value is evaluated immediately, but the pointed-to value is read later.
