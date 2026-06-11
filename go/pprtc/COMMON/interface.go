package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return (r.length * r.breadth)
}

func (c circle) area() float64 {
	return (math.Pi * c.radius * c.radius)
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func (c circle) perimeter() float64 {
	return (2 * math.Pi * c.radius)
}

func measure(shape string, g geometry) {
	fmt.Println("Area of "+shape+": ", g.area())
	fmt.Println("Perimeter of "+shape+": ", g.perimeter())
}

func main() {
	r := rectangle{length: 10.5, breadth: 4.2}
	c := circle{radius: 5.0}
	measure("Rectangle", r)
	measure("Circle", c)

}
