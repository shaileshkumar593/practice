package maain

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	Side float64
}

type Rectangle struct {
	length, width int
}

type Circle struct {
	radius int
}

func (c circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func calculateArea(shape interface{}) {
	switch s := shape.(type) {
	case Circle:
		fmt.Printf("Circle area: %.2f\n", s.area())
	case Rectangle:
		fmt.Printf("Rectangle area: %.2f\n", s.area())
	default:
		fmt.Println("Unknown shape")
	}
}

func main() {
	c := Circle{radius: 5}
	r := Rectangle{length: 4, width: 3}

	calculateArea(c)
	calculateArea(r)
}
