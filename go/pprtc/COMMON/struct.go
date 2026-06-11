package main

import (
	"fmt"
)

type Student struct {
	name   string
	m1, m2 float64
	rollno int
}

func main() {
	var s1 *Student = new(Student)
	s1.name = "shailesh"
	s1.m1 = 25.6
	s1.m2 = 82.9
	s1.rollno = 1

	fmt.Println(s1)
	fmt.Println(&s1)
}
