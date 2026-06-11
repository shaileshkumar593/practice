package main

import (
	"fmt"
)

type Rectangle struct {
	length  int
	breadth int
}

type Circle struct {
	radius int
}

type Service interface {
	area() int
}

func (rect Rectangle) area() int {
	return rect.length * rect.breadth
}

func (c Circle) area() int {
	return c.radius * c.radius
}

func main() {
	fmt.Println("Hello, World!")
	var rect = Rectangle{
		length:  10,
		breadth: 20,
	}
	var circle = Circle{
		radius: 10,
	}
	a1 := rect.area()
	b1 := circle.area()

	fmt.Println(a1)
	fmt.Print(b1)
}
