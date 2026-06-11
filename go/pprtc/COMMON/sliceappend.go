package main

import (
	"fmt"
)

func surprise(s []int, p []int) {
	s = append(s, 69)
	p1 := append(s, 25)
	p = append(p, 29)

	fmt.Println("in surprise ")
	fmt.Println("s", &s, s)
	fmt.Println("p", &p, p)
	fmt.Println("p1", &p1, p1)
}

func main() {
	var m = []int{1, 2, 3, 4}
	var p = make([]int, 5)
	fmt.Println("Before storing val --p", &p, p, len(p))
	for i := 0; i < len(p); i++ {
		fmt.Println("---i", i)
		fmt.Scanf("%d", &p[i])
		fmt.Println("i+++", i)

	}
	fmt.Println("After storing val ---p", &p, p)
	fmt.Println("Before  ---m", &m, m)
	m = append(m, 20)
	fmt.Println("After  ---m ", &m, m)
	surprise(m, p)

}
