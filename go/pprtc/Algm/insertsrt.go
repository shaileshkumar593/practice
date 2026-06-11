package main

import "fmt"

func main() {
	a := []int{5, 3, 3, 7, 9}
	for i := 1; i <= len(a)-1; i++ {
		key := a[i]
		j := i - 1

		for j >= 0 && key <= a[j] {
			a[j+1] = a[j]
			j = j - 1

		}
		a[j+1] = key
	}
	fmt.Println(a)
}
