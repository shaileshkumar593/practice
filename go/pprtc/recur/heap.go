package main

import "fmt"

func main() {
	f := test()
	fmt.Println(f())
}

func test() func() int {
	x := 5
	return func() int {
		return x
	}
}
