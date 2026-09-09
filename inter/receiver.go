package main

import "fmt"

// Struct type
type User struct {
	Name string
	Age  int
}

// New defined type whose underlying type is int
type MyIntss int

// Method with struct receiver
func (u User) PrintFunc() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
}

// Method with MyInt receiver
func (n MyIntss) PrintFunc() {
	fmt.Println("Value:", n)
}

func main() {
	// Struct value
	u := User{
		Name: "Shailesh",
		Age:  30,
	}

	// MyInt value
	n := MyIntss(100)

	// Call methods
	u.PrintFunc()
	n.PrintFunc()
}

/*

A Go method receiver must be a defined type declared in the same package.
any cannot be used as a receiver because any is an alias for interface{},
not a new defined type. The same reason applies to built-in types like int and string.
If I need a method, I create a defined type such as type MyInt int. If I need to accept
values of any type, I use any as a function parameter or interface parameter instead." */

/*
You cannot define two methods with the same name on the same receiver type:

type User struct {
	Name string
}

func (u User) PrintFunc() {
}

func (u User) PrintFunc() { // ❌ compile error
}
Go does not support method overloading.


Can we define methods on int, string, or other built-in types?

No. You define a new local type whose underlying type is the built-in type, and attach methods to that new type.
*/
