package main

import (
	"fmt"
)

// ✅ Interface declaration
type Shape interface {
	Area() float64
}

// ✅ Struct declaration
type Rectangle struct {
	Width  float64
	Height float64
}

// Implementing the Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {

	// ✅ Array declaration
	var numbers [3]int = [3]int{10, 20, 30}
	fmt.Println("Array:", numbers)

	// ✅ Slice declaration
	fruits := []string{"Apple", "Banana", "Cherry"}
	fruits = append(fruits, "Mango")
	fmt.Println("Slice:", fruits)

	// ✅ Map declaration
	countryCapital := map[string]string{
		"India":  "New Delhi",
		"Japan":  "Tokyo",
		"France": "Paris",
	}
	countryCapital["Germany"] = "Berlin"
	fmt.Println("Map:", countryCapital)

	// ✅ Struct usage
	rect := Rectangle{Width: 10.5, Height: 5.5}
	fmt.Printf("Struct: %+v\n", rect)

	// ✅ Interface usage
	var s Shape
	s = rect
	fmt.Println("Interface (Area):", s.Area())
}
