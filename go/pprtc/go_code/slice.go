package main

import "fmt"

func main() {
	a := make([]int, 4)
	fmt.Println(a)
	var i int = 0
	for i = 0; i < len(a); i++ {
		a[i] = i
	}
	append(a, 25)
	fmt.Println(a)
}
