package main

import "fmt"

func main() {
	arr := [4]int{4, 7, 8}
	a := [...]int{5, 9}

	var arr1 [3]int
	arr1 = arr
	fmt.Println("Hello")
	sum := 0
	for i := 0; i < 4; i++ {
		sum = sum + arr[i]
		fmt.Println("Sum ", sum)
	}

}
