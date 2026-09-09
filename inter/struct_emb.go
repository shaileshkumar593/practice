package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Person struct {
	Name string
	Age  int

	Address // Embedded struct
}

func main() {
	p := Person{
		Name: "Shailesh",
		Age:  30,
		Address: Address{
			City:  "Ranchi",
			State: "Jharkhand",
		},
	}

	fmt.Println(p.Name)
	fmt.Println(p.City)         // promoted field
	fmt.Println(p.Address.City) // explicit access
}
