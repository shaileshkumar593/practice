package main

import "fmt"

// =====================================================
// 1. STRUCT
// =====================================================

type Employee struct {
	Name   string
	Salary int
}

// Value receiver
// It receives a copy of Employee.
func (e Employee) GetSalary() int {
	return e.Salary
}

// Value receiver
// It only reads the Employee data.
func (e Employee) PrintDetails() {
	fmt.Printf("Name: %s, Salary: %d\n", e.Name, e.Salary)
}

// Pointer receiver
// It modifies the original Employee.
func (e *Employee) IncreaseSalary(amount int) {
	e.Salary += amount
}

// Pointer receiver
func (e *Employee) ChangeName(name string) {
	e.Name = name
}

// =====================================================
// 2. CUSTOM TYPES
// =====================================================

// New defined type.
// EmployeeID is NOT the same type as int.
type EmployeeID int

// Method on EmployeeID.
func (id EmployeeID) IsValid() bool {
	return id > 0
}

// Method on EmployeeID.
func (id EmployeeID) Print() {
	fmt.Println("Employee ID:", id)
}

// =====================================================
// 3. TYPE ALIAS
// =====================================================

// EmployeeCount is an alias for int.
// EmployeeCount and int are exactly the same type.
type EmployeeCount = int

// =====================================================
// 4. INTERFACE
// =====================================================

type Printable interface {
	PrintDetails()
}

// Employee implements Printable because Employee
// has PrintDetails().
func PrintEmployee(p Printable) {
	p.PrintDetails()
}

// =====================================================
// 5. MAIN
// =====================================================

func main() {

	// -------------------------------------------------
	// Struct creation
	// -------------------------------------------------

	employee := Employee{
		Name:   "Shailesh",
		Salary: 50000,
	}

	fmt.Println("Initial employee:")
	employee.PrintDetails()

	// -------------------------------------------------
	// Value receiver
	// -------------------------------------------------

	salary := employee.GetSalary()

	fmt.Println("Salary:", salary)

	// -------------------------------------------------
	// Pointer receiver
	// -------------------------------------------------

	employee.IncreaseSalary(10000)

	fmt.Println("\nAfter salary increase:")
	employee.PrintDetails()

	employee.ChangeName("Rahul")

	fmt.Println("\nAfter name change:")
	employee.PrintDetails()

	// -------------------------------------------------
	// Custom type
	// -------------------------------------------------

	var id EmployeeID = 101

	fmt.Println("\nEmployee ID:")
	id.Print()

	fmt.Println("Is ID valid?", id.IsValid())

	// -------------------------------------------------
	// Type safety
	// -------------------------------------------------

	var normalInt int = 100

	var employeeID EmployeeID = 200

	fmt.Println("\nNormal int:", normalInt)
	fmt.Println("EmployeeID:", employeeID)

	// This is NOT allowed:
	//
	// employeeID = normalInt
	//
	// Compile error:
	// cannot use normalInt (variable of type int)
	// as EmployeeID value

	// Explicit conversion is required:
	employeeID = EmployeeID(normalInt)

	fmt.Println("After conversion:", employeeID)

	// -------------------------------------------------
	// Type alias
	// -------------------------------------------------

	var count EmployeeCount = 10

	var number int = 20

	// EmployeeCount is an alias for int,
	// therefore assignment is allowed.
	count = number

	fmt.Println("\nEmployee count:", count)

	// -------------------------------------------------
	// Interface
	// -------------------------------------------------

	fmt.Println("\nInterface example:")

	var printable Printable = employee

	printable.PrintDetails()

	// Or directly:
	PrintEmployee(employee)

	// -------------------------------------------------
	// Pointer to struct
	// -------------------------------------------------

	employeePtr := &Employee{
		Name:   "Amit",
		Salary: 60000,
	}

	employeePtr.IncreaseSalary(5000)

	fmt.Println("\nPointer employee:")
	employeePtr.PrintDetails()
}
