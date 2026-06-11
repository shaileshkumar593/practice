package main

import "fmt"

// defining strtucture
// type nameoftype struct

type Address struct {
	name    string
	city    string
	pincode int
}

func main() {
	// declaring variable of type address
	var a Address
	fmt.Println(a)
	// Address variable a is initialize with zero value
	a1 := Address{
		"Shailesh",
		"Delhi",
		110016,
	}

	a2 := Address{
		name:    "Priyanka",
		city:    "Gujrat",
		pincode: 51600,
	}

	a3 := Address{name: "Rakesh"}

	fmt.Println("Address is :", a)
	fmt.Println("Address of Shailesh is :", a1)
	fmt.Println("Address of Priyanka is :", a2)
	fmt.Println("Address is :", a3)
}
