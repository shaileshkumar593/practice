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

		if n == 3 { // output till 3 then escape 4,5  O(1) memory use 
			// good for generating billions of memory
			break
		}
	}
}