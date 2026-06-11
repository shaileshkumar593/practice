package main

import "fmt"

func main() {
	var a = []int{5, 9, 7, 6, 3}

	for i, j := 0, len(a)-1; i <= j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println(a)
}
