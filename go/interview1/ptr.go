package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 1. Pass by value: original does NOT change
func updateValue(x int) {
	x = 100
}

// 2. Pass by pointer: original changes
func updatePointer(x *int) {
	*x = 100
}

// 3. Struct pointer
func updateUser(u *User) {
	u.Name = "Updated"
	u.Age = 99
}

// 4. Pointer receiver
func (u *User) Birthday() {
	u.Age++
}

func main() {
	fmt.Println("CASE 1: zero value of pointer")
	var p *int
	fmt.Println(p) // nil

	// fmt.Println(*p) // panic: nil pointer dereference

	fmt.Println("\nCASE 2: pointer using &")
	x := 10
	ptr := &x
	fmt.Println("x =", x)
	fmt.Println("ptr address =", ptr)
	fmt.Println("*ptr value =", *ptr)

	fmt.Println("\nCASE 3: modify value using pointer")
	*ptr = 20
	fmt.Println("x after *ptr=20:", x)

	fmt.Println("\nCASE 4: new(int)")
	n := new(int)
	fmt.Println("n address =", n)
	fmt.Println("*n default value =", *n)

	*n = 50
	fmt.Println("*n after update =", *n)

	fmt.Println("\nCASE 5: pass by value")
	a := 10
	updateValue(a)
	fmt.Println("a after updateValue:", a) // 10

	fmt.Println("\nCASE 6: pass by pointer")
	updatePointer(&a)
	fmt.Println("a after updatePointer:", a) // 100

	fmt.Println("\nCASE 7: struct value")
	u1 := User{Name: "Shailesh", Age: 30}
	fmt.Println("before:", u1)

	updateUser(&u1)
	fmt.Println("after:", u1)

	fmt.Println("\nCASE 8: struct pointer using new")
	u2 := new(User)
	fmt.Println("default user:", *u2)

	u2.Name = "Go Developer"
	u2.Age = 35
	fmt.Println("updated user:", *u2)

	fmt.Println("\nCASE 9: struct pointer using &User{}")
	u3 := &User{Name: "Backend Engineer", Age: 40}
	fmt.Println("u3:", *u3)

	u3.Birthday()
	fmt.Println("after birthday:", *u3)

	fmt.Println("\nCASE 10: nil pointer check")
	var user *User

	if user == nil {
		fmt.Println("user is nil")
	}

	user = &User{Name: "Active User", Age: 25}

	if user != nil {
		fmt.Println("user is not nil:", *user)
	}
}
