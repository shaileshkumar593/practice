package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof")
}

func main() {
	var s1 Speaker = Dog{}
	var s2 Speaker = &Dog{}

	s1.Speak()
	s2.Speak()
}
