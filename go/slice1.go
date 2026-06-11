package main

import "fmt"

func main() {
	a := []int{3, 5, 6, 7, 8, 9}

	s1 := a[1:2]
	s2 := a[3:5]

	fmt.Println(a, " : ", s1, " : ", s2)

	s2 = append(s2, 44, 33, 77)
	fmt.Println(a, " : ", s1, " : ", s2)
}
