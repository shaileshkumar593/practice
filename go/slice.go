package main

import "fmt"

func main() {
	a := []int{2, 4, 7, 8, 9}

	s1 := a[1:2]

	s2 := a[3:4]

	fmt.Println(s1, s2, len(s1), cap(s1), len(s2), cap(s2))
	s2 = append(s2, 11)

	fmt.Print(a, " : ", s1, " : ", s2)

	s2 = append(s2, 55)
	fmt.Print(a, " : ", s1, " : ", s2)
}
