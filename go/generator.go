package main

import "fmt"

func Numbers(yield func(int) bool) {

	for i := 1; i <= 5; i++ {

		v := yield(i)
		fmt.Println(v)
		if !v {
			fmt.Println("hi")
			return
		}
	}
}

func main() {

	for n := range Numbers {
		fmt.Println(n)
	}
}