package main

import "fmt"

//Left Even and Right odd
func main() {
	var x = []int{1, 2, 3, 4, 5, 6}
	var l int = 0
	var h int = len(x) - 1
	fmt.Println("Array before sorting")
	fmt.Println(x)
	for l <= h {
		if x[l]%2 != 0 {
			if x[h]%2 == 0 {
				x[l], x[h] = x[h], x[l]
				l += 1
				h -= 1
			} else {
				h -= 1
			}
		} else {
			l += 1
		}
	}
	fmt.Println("Array after sorting")
	fmt.Println(x)
}
