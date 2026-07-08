package main

import (
	"fmt"
)

func main() {
	a := []int{8, 9, 7}

	for i := range a {
		fmt.Println(i)
	}
}
// 0,1,2  access index  by default