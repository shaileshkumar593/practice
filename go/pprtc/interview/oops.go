package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
}

// Define a Circle struct
type Circle struct {
	Radius float64
}

// Implement the Area method for Circle using a receiver function
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Define a Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// A function that accepts any type satisfying the Shape interface
func PrintArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}

	// Both Circle and Rectangle can be used with PrintArea
	PrintArea(c)
	PrintArea(r)
}

/*
Key Object-Oriented Features and Implementation
Feature 			        Go Implementation	               Description
Classes and Objects	    Structs and Receiver Functions	    Go uses struct types to group data fields, similar to classes in other languages. Behavior is added using methods (functions with a special receiver argument that binds the function to the struct type).
Encapsulation		   Exported/Unexported Identifiers	    Access control is managed by naming conventions at the package level. An identifier (field, method, function, or type) starting with a capital letter is exported (public), while one starting with a lowercase letter is unexported (private to the package).
Inheritance	             Composition and Embedding	        Go does not support class-based inheritance. Instead, it favors composition, where a struct can embed another struct (anonymously). This allows the outer struct to directly access the fields and methods of the embedded type, providing a flexible alternative to traditional inheritance.
Polymorphism and Abstraction	Interfaces	                Go's interfaces define a set of method signatures, not implementation details. Any type that implements all the methods of an interface implicitly satisfies that interface. This allows different concrete types to be treated uniformly based on their behavior, enabling polymorphism and abstraction without an explicit implements keyword.

*/
