package main

import (
	"fmt"
)

func main() {
	p := [6]int{3, 8, 2, 99, 24, 44}

	s := p[2:4]

	fmt.Println(s)

	/*
		len = j - i
		cap = original_capacity - i

		s[start:end:max]

		len = end - start
		= 3 - 1
		= 2

		cap = max - start
			= 3 - 1
			= 2
	*/
	fmt.Println(p)
	fmt.Println(s)

	s = append(s, 233) // modify p
	fmt.Println(p)
	fmt.Println(s)

	fmt.Println("::::::: g::::::::")

	g := p[1:6]

	fmt.Println(p)
	fmt.Println(g)
	g = append(g, 242) // create new slice as it reaches capacity
	fmt.Println(p)
	fmt.Println(g)

	//h :=  p[1:5:2] // nvalid slice indices: 2 < 5

	fmt.Println("::::::: h::::::::")
	h := p[1:3:3]
	fmt.Println(p)
	fmt.Println(h)

	h = append(h, 55, 66)
	fmt.Println(p)
	fmt.Println(h) // based on above expression cap exceed so new slice is allocated

}
