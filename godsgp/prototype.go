package main

import "fmt"

type Address struct {
	City string
}

type User struct {
	Name    string
	Age     int
	Address *Address
}

func (u *User) Clone() *User {
	return &User{
		Name: u.Name,
		Age:  u.Age,
		Address: &Address{
			City: u.Address.City,
		},
	}
}

func main() {
	user1 := &User{
		Name: "John",
		Age:  30,
		Address: &Address{
			City: "Delhi",
		},
	}

	user2 := user1.Clone()

	user2.Name = "Mike"
	user2.Address.City = "Mumbai"

	fmt.Printf("User1: %+v\n", user1)
	fmt.Printf("User2: %+v\n", user2)
}
