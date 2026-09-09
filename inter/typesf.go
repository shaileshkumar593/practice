package main

import "fmt"

// Domain-specific types
type UserID int64
type OrderID int64

// Function specifically expects UserID
func GetUser(id UserID) {
	fmt.Println("User:", id)
}

func main() {

	// Correct
	userID := UserID(101)

	GetUser(userID) // ✅

	// Different defined type
	orderID := OrderID(500)

	// GetUser(orderID) // ❌ compile error

	// Explicit conversion is required
	GetUser(UserID(orderID)) // ✅
}

/*
Go's compiler generates the error because UserID and OrderID are distinct defined types.
Their underlying type is int64, but Go's type system does not consider two separately
defined types interchangeable. This provides compile-time type safety and prevents accidentally
 passing an OrderID where a UserID is expected.
*/
