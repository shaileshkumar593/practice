package main

import "fmt"

/*

	Sure! In Go, the ++ and -- operators can only be applied to numeric types
	(int, int8, int16, int32, int64, uint, float32, float64, etc.). They cannot
	be used on slices, strings, booleans, structs, or maps.

*/

func main() {
	// Signed integers
	var a int = 5
	a++
	fmt.Println("int a after ++:", a)
	a--
	fmt.Println("int a after --:", a)

	var b int8 = 10
	b++
	fmt.Println("int8 b after ++:", b)
	b--
	fmt.Println("int8 b after --:", b)

	var c int16 = 100
	c++
	fmt.Println("int16 c after ++:", c)
	c--
	fmt.Println("int16 c after --:", c)

	var d int32 = 1000
	d++
	fmt.Println("int32 d after ++:", d)
	d--
	fmt.Println("int32 d after --:", d)

	var e int64 = 10000
	e++
	fmt.Println("int64 e after ++:", e)
	e--
	fmt.Println("int64 e after --:", e)

	// Unsigned integers
	var u uint = 5
	u++
	fmt.Println("uint u after ++:", u)
	u--
	fmt.Println("uint u after --:", u)

	var u8 uint8 = 255
	u8++
	fmt.Println("uint8 u8 after ++:", u8) // wraps around to 0
	u8--
	fmt.Println("uint8 u8 after --:", u8)

	var u16 uint16 = 65535
	u16++
	fmt.Println("uint16 u16 after ++:", u16) // wraps around to 0
	u16--
	fmt.Println("uint16 u16 after --:", u16)

	var u32 uint32 = 1000
	u32++
	fmt.Println("uint32 u32 after ++:", u32)
	u32--
	fmt.Println("uint32 u32 after --:", u32)

	var u64 uint64 = 10000
	u64++
	fmt.Println("uint64 u64 after ++:", u64)
	u64--
	fmt.Println("uint64 u64 after --:", u64)

	// Float types
	var f32 float32 = 1.5
	f32++
	fmt.Println("float32 f32 after ++:", f32)
	f32--
	fmt.Println("float32 f32 after --:", f32)

	var f64 float64 = 2.5
	f64++
	fmt.Println("float64 f64 after ++:", f64)
	f64--
	fmt.Println("float64 f64 after --:", f64)

	//x := 5
	//y := x++  // invalid
}

/*


	x := 5
	y := x++   // ❌ INVALID in Go

	++ and -- are statements, not expressions.

	This means they change the variable, but do not produce a value.

	x := 5
	x++      // ✅ valid, increments x by 1
	fmt.Println(x) // prints 6

	y := x++  // ❌ INVALID in Go

	Why invalid?

	x++ does not return a value.

	:= expects a value to assign to y.

	In C/C++/Java, x++ is an expression that returns the old value of x, so y := x++ works.

	In Go, ++ and -- are only statements, not expressions.

	How to do similar in Go
	If you want y to get the old value of x and then increment x, you have to split it into two statements:
	x := 5
	y := x   // assign old value first
	x++      // increment x
	fmt.Println("x =", x) // 6
	fmt.Println("y =", y) // 5

	In Go, ++ and -- are statements only, so they cannot be part of expressions or assignments.


	| Feature / Operation                   | C / C++ / Java                                                        | Go                                                                       |
| ------------------------------------- | --------------------------------------------------------------------- | ------------------------------------------------------------------------ |
| `x++` (post-increment)                | **Expression**: returns the current value of `x`, then increments `x` | **Statement only**: increments `x`, but returns **no value**             |
| `++x` (pre-increment)                 | **Expression**: increments `x`, then returns the new value            | **Statement only**: increments `x`, but returns **no value**             |
| `y := x++`                            | ✅ Valid. `y` gets old value of `x`; `x` is incremented                | ❌ Invalid. `x++` is not an expression; cannot assign to `y`              |
| `y := ++x`                            | ✅ Valid. `x` is incremented first; `y` gets new value                 | ❌ Invalid. `++x` is a statement, not an expression                       |
| Usage in loops                        | `for (int i=0;i<10;i++)` ✅ Valid                                      | `for i:=0; i<10; i++` ✅ Valid, because `i++` is **standalone statement** |
| Increment in expressions              | `z = x++ + 5` ✅ Valid                                                 | ❌ Invalid. `x++` cannot be used inside expressions                       |
| Correct Go alternative for `y := x++` | N/A                                                                   | `go y := x; x++ `                                                        |


*/
