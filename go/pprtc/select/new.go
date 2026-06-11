package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    // Using new() to allocate memory for a Person struct
    p := new(Person)

    fmt.Printf("%T\n", p)

    // Accessing struct fields using the pointer
    p.Name = "Alice"
    p.Age = 30

    // Displaying the values
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
}