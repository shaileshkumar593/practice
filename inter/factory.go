package main

import "fmt"

// ============================================================
// Interface
// ============================================================

type PrintMethod interface {
	PrintFunc()
}

// ============================================================
// Employee
// ============================================================

type Employee struct {
	Name   string
	Salary int
}

// Employee implements PrintFunc()
func (e Employee) PrintFunc() {
	fmt.Println("Employee Details")
	fmt.Println("Name:", e.Name)
	fmt.Println("Salary:", e.Salary)
}

// ============================================================
// User
// ============================================================

type User struct {
	ID   int
	Name string
}

// User also implements PrintFunc()
func (u User) PrintFunc() {
	fmt.Println("User Details")
	fmt.Println("ID:", u.ID)
	fmt.Println("Name:", u.Name)
}

// ============================================================
// printInfo accepts anything that satisfies PrintMethod
// ============================================================

func printInfo(p PrintMethod) {
	p.PrintFunc()
}

// ============================================================
// Main
// ============================================================

func main() {

	employee := Employee{
		Name:   "Shailesh",
		Salary: 100000,
	}

	user := User{
		ID:   101,
		Name: "Rahul",
	}

	// Employee implements PrintMethod
	printInfo(employee)

	fmt.Println()

	// User implements PrintMethod
	printInfo(user)
}
