package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {

	// making pointer to structure
	emp8 := &Employee{"Shailesh", "Kumar", 32, 40000}

	fmt.Println("First Name :", (*emp8).firstName)

	//emp8.lastName instead of explicit deference (*emp8).firstName
	fmt.Println("Last Name :", emp8.lastName)
	fmt.Println("Age :", (*emp8).age)
	fmt.Println("salary is :", emp8.salary)
}
