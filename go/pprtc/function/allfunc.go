package main

import (
	"fmt"
)

// ---------- 13. Struct + Methods ----------
type User struct {
	Name string
	Age  int
}

// Value receiver
func (u User) Print() {
	fmt.Println("Value Receiver:", u.Name, u.Age)
}

// Pointer receiver
func (u *User) Rename(name string) {
	u.Name = name
}

// ---------- 15. Function Type ----------
type Operation func(int, int) int

// ---------- 1. Basic Function ----------
func add(a, b int) int {
	return a + b
}

// ---------- 3. Multiple Return ----------
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
	}
	return a / b, nil
}

// ---------- 4. Named Return ----------
func calc(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return
}

// ---------- 5. Variadic ----------
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// ---------- 8. Function as argument ----------
func operate(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

// ---------- 9. Return function ----------
func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// ---------- 10. Closure ----------
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// ---------- 17. Recursion ----------
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// ---------- 12. Panic & Recover ----------
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	if b == 0 {
		panic("divide by zero")
	}

	fmt.Println("Division:", a/b)
}

// ---------- 16. Higher-order ----------
func apply(nums []int, fn func(int) int) []int {
	result := make([]int, 0, len(nums))
	for _, n := range nums {
		result = append(result, fn(n))
	}
	return result
}

func main() {

	// 1. Basic
	fmt.Println("Add:", add(2, 3))

	// 3. Multiple return
	res, err := divide(10, 2)
	fmt.Println("Divide:", res, err)

	// 4. Named return
	s, p := calc(3, 4)
	fmt.Println("Calc:", s, p)

	// 5. Variadic
	fmt.Println("Sum:", sum(1, 2, 3, 4))

	// 6. Function as value
	fn := func(a, b int) int {
		return a - b
	}
	fmt.Println("Function as value:", fn(10, 5))

	// 7. Anonymous
	func() {
		fmt.Println("Anonymous function")
	}()

	// 8. Function as argument
	fmt.Println("Operate:", operate(5, 2, add))

	// 9. Return function
	double := multiplier(2)
	fmt.Println("Multiplier:", double(10))

	// 10. Closure
	c := counter()
	fmt.Println("Counter:", c(), c(), c())

	// 11. Defer
	func() {
		defer fmt.Println("Defer executed")
		fmt.Println("Inside function")
	}()

	// 12. Panic & recover
	safeDivide(10, 0)

	// 13 + 14. Struct methods
	u := User{Name: "Shailesh", Age: 30}
	u.Print()
	u.Rename("GoDev")
	fmt.Println("After Rename:", u)

	// 15. Function type
	var op Operation = add
	fmt.Println("Operation:", op(10, 20))

	// 16. Higher-order
	arr := []int{1, 2, 3}
	result := apply(arr, func(x int) int {
		return x * x
	})
	fmt.Println("Apply:", result)

	// 17. Recursion
	fmt.Println("Factorial:", factorial(5))

	// 18. Goroutine function
	done := make(chan bool)
	go func() {
		fmt.Println("Goroutine running")
		done <- true
	}()
	<-done

	// 19. Closure bug fix
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println("Correct closure:", i)
		}(i)
	}
}
