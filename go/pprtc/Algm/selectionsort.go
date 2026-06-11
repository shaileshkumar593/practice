package main

import "fmt"

func main() {
	a := []int{5, 3, 3, 7, 9}
	for i := 0; i < len(a)-1; i++ {
		min := i

		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[min], a[i] = a[i], a[min]
	}

	fmt.Println(a)
}
