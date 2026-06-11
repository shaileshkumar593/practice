package main

import "fmt"

func main() {

	// =========================
	// CASE 1: Simple defer with parameter
	// =========================
	x := 10

	defer fmt.Println("Case 1:", x)
	// WHY: argument x is evaluated IMMEDIATELY at defer time (x=10 stored)
	/*
		fmt.Println(x) is a function call

		Go evaluates arguments immediately at defer time
	*/

	x = 20
	// WHY NOT 20? because defer already captured 10

	// =========================
	// CASE 2: Defer with function parameter (explicit)
	// =========================
	y := 10

	defer func(a int) {
		fmt.Println("Case 2:", a)
	}(y)
	// WHY: y is passed as argument immediately → a = 10 stored

	y = 20
	// WHY NOT 20? because value already copied into parameter

	// =========================
	// CASE 3: Closure defer (NO parameter)
	// =========================
	z := 10

	defer func() {
		fmt.Println("Case 3:", z)
	}()
	// WHY: closure captures variable reference, NOT value

	z = 20
	// WHY: defer reads latest value → 20 printed

	// =========================
	// CASE 4: Difference between parameter vs closure
	// =========================
	a := 1

	defer func(val int) {
		fmt.Println("Case 4 (param):", val)
	}(a)
	// WHY: parameter captured immediately → val = 1

	defer func() {
		fmt.Println("Case 4 (closure):", a)
	}()
	// WHY: closure reads latest value

	a = 100
	// WHY:
	// param version prints 1
	// closure version prints 100

	// =========================
	// CASE 5: Expression inside defer argument
	// =========================
	b := 5

	defer fmt.Println("Case 5:", b+10)
	// WHY: expression b+10 is evaluated immediately → 15 stored

	b = 100
	// WHY NOT 110? because result already computed

	// =========================
	// CASE 6: Multiple defer + parameter behavior
	// =========================
	c := 1

	defer func(val int) {
		fmt.Println("Case 6-1:", val)
	}(c)
	// captures 1

	c = 2

	defer func(val int) {
		fmt.Println("Case 6-2:", val)
	}(c)
	// captures 2

	c = 3

	// =========================
	// EXECUTION ORDER NOTE:
	// =========================
	// WHY OUTPUT ORDER?
	// - defer stack = LIFO (last in, first out)
	// - BUT parameter values already fixed at defer time

	/*
		EXPECTED OUTPUT ORDER:
		Case 6-2: 2
		Case 6-1: 1
		Case 4 (closure): 100
		Case 4 (param): 1
		Case 3: 20
		Case 2: 10
		Case 1: 10
	*/

	fmt.Println("Main finished")
}
