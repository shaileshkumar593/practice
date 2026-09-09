package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ============================================================
// 1. Domain-specific ID types
// ============================================================

// UserID is a distinct type based on int64.
type UserID1 int64

// OrderID is another distinct type based on int64.
// UserID and OrderID are NOT interchangeable.
type OrderID1 int64

// ============================================================
// 2. Domain-specific Email type
// ============================================================

// Email is a distinct type based on string.
type Email string

// IsValid validates the email.
func (e Email) IsValid() bool {
	return strings.Contains(string(e), "@")
}

// ============================================================
// 3. User domain model
// ============================================================

type User1 struct {
	ID    UserID `json:"id"`
	Email Email  `json:"email"`
	Name  string `json:"name"`
}

// ============================================================
// 4. User validation
// ============================================================

func (u User1) IsValid() bool {
	if u.ID <= 0 {
		return false
	}

	if !u.Email.IsValid() {
		return false
	}

	if strings.TrimSpace(u.Name) == "" {
		return false
	}

	return true
}

// ============================================================
// 5. Function expecting UserID
// ============================================================

func GetUser1(id UserID1) {
	fmt.Println("Fetching user with ID:", id)
}

// ============================================================
// 6. Function expecting OrderID
// ============================================================

func GetOrder(id OrderID1) {
	fmt.Println("Fetching order with ID:", id)
}

// ============================================================
// 7. Main
// ============================================================

func main() {

	// --------------------------------------------------------
	// UserID
	// --------------------------------------------------------

	userID1 := UserID1(101)

	fmt.Println("User ID:", userID1)

	GetUser(userID1) // ✅ Correct

	// --------------------------------------------------------
	// OrderID
	// --------------------------------------------------------

	orderID := OrderID1(500)

	fmt.Println("Order ID:", orderID)

	GetOrder(orderID) // ✅ Correct

	// --------------------------------------------------------
	// Type safety
	// --------------------------------------------------------

	// GetUser(orderID)
	//
	// ❌ Compile error:
	//
	// cannot use orderID (variable of type OrderID)
	// as UserID value in argument to GetUser
	//
	// Why?
	//
	// UserID and OrderID are different defined types.

	// --------------------------------------------------------
	// Explicit conversion
	// --------------------------------------------------------

	GetUser(UserID1(orderID)) // ✅ Explicit conversion

	fmt.Println()

	// --------------------------------------------------------
	// Email domain type
	// --------------------------------------------------------

	email := Email("shailesh@example.com")

	fmt.Println("Email:", email)

	if email.IsValid() {
		fmt.Println("Email is valid")
	} else {
		fmt.Println("Email is invalid")
	}

	fmt.Println()

	// --------------------------------------------------------
	// Invalid email
	// --------------------------------------------------------

	invalidEmail := Email("shailesh-example.com")

	fmt.Println("Email:", invalidEmail)

	if invalidEmail.IsValid() {
		fmt.Println("Email is valid")
	} else {
		fmt.Println("Email is invalid")
	}

	fmt.Println()

	// --------------------------------------------------------
	// Create User
	// --------------------------------------------------------

	user := User1{
		ID:    UserID(101),
		Email: Email("shailesh@example.com"),
		Name:  "Shailesh",
	}

	fmt.Println("User:")
	fmt.Println("ID:", user.ID)
	fmt.Println("Email:", user.Email)
	fmt.Println("Name:", user.Name)

	// --------------------------------------------------------
	// Validate User
	// --------------------------------------------------------

	if user.IsValid() {
		fmt.Println("User is valid")
	} else {
		fmt.Println("User is invalid")
	}

	fmt.Println()

	// --------------------------------------------------------
	// JSON
	// --------------------------------------------------------

	data, err := json.MarshalIndent(user, "", "  ")

	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	fmt.Println("JSON:")
	fmt.Println(string(data))
}
