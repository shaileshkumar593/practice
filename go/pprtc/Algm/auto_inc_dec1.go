package main

import (
	"fmt"
)

func main() {
	x := 5

	fmt.Println(" x ", x)

	x++
	fmt.Println(" x ", x)
	// fmt.Println(" x ", x++) error

}

/*

| Code               | Error                     | Explanation                                                                    |
| ------------------ | ------------------------- | ------------------------------------------------------------------------------ |
| `y := x++`         | `cannot use x++ as value` | `x++` is a statement, not an expression. Go cannot assign its "result" to `y`. |
| `y := ++x`         | `cannot use ++x as value` | Same reason: `++x` does not return a value.                                    |
| `z := 5 + x++`     | `cannot use x++ as value` | Expressions cannot contain `++`/`--`.                                          |
| `fmt.Println(x++)` | `cannot use x++ as value` | `++` cannot be used as a function argument.                                    |
| `arr[x++] = 10`    | `cannot use x++ as value` | Index expressions cannot include `++`.                                         |
*/

/*

	Go ++ and --: Complete Reference
1️⃣ Simple Increment / Decrement
x := 5
x++  // ✅ valid, x becomes 6
x--  // ✅ valid, x becomes 5
Why: standalone statements, numeric type.
Why not: N/A.

2️⃣ Pre/Post Assignment (Invalid)
y := x++  // ❌ invalid
y := ++x  // ❌ invalid
Why not: ++ / -- are statements, not expressions. They don’t return a value.

Alternative:

y := x
x++
3️⃣ Expressions (Invalid)
z := x++ + 5  // ❌ invalid
fmt.Println(x++) // ❌ invalid
Why not: Cannot use ++ / -- inside expressions or function calls.

Alternative: split statements:

z := x + 5
x++
fmt.Println(z)
4️⃣ Loop Increment / Decrement (Valid)
for i := 0; i < 5; i++ { fmt.Println(i) } // ✅ valid
for i := 5; i > 0; i-- { fmt.Println(i) } // ✅ valid
Why: Go allows ++/-- as standalone statements at the loop iteration.

5️⃣ On Different Integer Types (Valid)
var a int8 = 10
var b int16 = 100
var c int64 = 10000
a++; b++; c++
Why: Works on all integer types.

6️⃣ On Unsigned Integers (Valid)
var u uint = 5
u++  // ✅ increments
u--  // ✅ decrements
Why: Works on uint, uint8, uint16, uint32, uint64.

Note: Underflow/overflow wraps around.

7️⃣ On Float Types (Valid)
var f float32 = 2.5
f++
f--
Why: Go allows ++ / -- on float types as statements.

8️⃣ On Non-numeric Types (Invalid)
s := "hello"
s++ // ❌ invalid
Why not: Go only allows numeric types. Strings, bools, arrays, maps, structs are invalid.

9️⃣ On Arrays / Slices / Maps (Invalid)
arr := []int{1,2}
arr++ // ❌ invalid
Why not: ++ / -- only work on individual numeric variables, not collections.

10️⃣ On Struct Fields (Valid if Numeric)
type Point struct { X int }
p := Point{X:5}
p.X++
Why: Field is numeric; statement modifies value.

11️⃣ On Pointer Dereference (Valid)
x := 5
px := &x
(*px)++  // ✅ increments x via pointer
Why: *px is numeric.

12️⃣ In Function Call (Invalid)
fmt.Println(x++) // ❌ invalid
Why not: x++ is statement; function argument requires expression.

Alternative:

fmt.Println(x)
x++
13️⃣ As Return Value (Invalid)
return x++ // ❌ invalid
Why not: statements cannot be returned as values.

Alternative: split:

val := x
x++
return val
14️⃣ As Map Index (Invalid)
m := map[int]int{}
m[x++] = 5 // ❌ invalid
Why not: index expressions need value; x++ is statement.

15️⃣ Using in Assignment (Invalid)
y := x-- // ❌ invalid
Why not: -- is statement; cannot assign as expression.

16️⃣ With Bitwise Operations (Invalid)
z := x++ & 3 // ❌ invalid
Why not: Cannot use ++ inside expression.

17️⃣ Inside Slice / Array Assignment (Invalid)
arr[x++] = 10 // ❌ invalid
Why not: ++ cannot appear in index or RHS.

Alternative:

arr[x] = 10
x++
18️⃣ In Conditional Expression (Invalid)
if x++ > 5 {} // ❌ invalid
Why not: statements cannot be in expressions.

Alternative:

cond := x > 5
x++
if cond {}
19️⃣ On Constants (Invalid)
const c = 5
c++ // ❌ invalid
Why not: constants are immutable.

20️⃣ On Multiple Variables at Once (Invalid)
a, b := 1, 2
a++, b++ // ❌ invalid
Why not: Go requires one statement per increment.

Alternative:

a++
b++
21️⃣ On Loop Variables (Valid)
for i := 0; i < 10; i++ { } // ✅ valid
Why: loop iteration supports ++ and -- as statements.

22️⃣ Using in Switch / Case (Invalid)
switch x++ { } // ❌ invalid
Why not: switch expression expects value, not statement.

23️⃣ Inside Goroutine or Channel (Valid)
x := 5
go func() { x++ }() // ✅ valid, increment inside function
Why: ++ is statement; can run anywhere statements are allowed.

24️⃣ Overflow / Underflow (Depends on Type)
var u8 uint8 = 255
u8++ // wraps to 0 ✅
u8-- // wraps to 255 ✅
Why: unsigned integers wrap around.

25️⃣ Incrementing Expressions (Invalid)
(x + 5)++ // ❌ invalid
Why not: must operate on addressable variable, not temporary value.

✅ Summary Rules
Valid: standalone numeric variables or dereferenced pointers.

Invalid: inside expressions, assignments, function arguments, array indices, constants, non-numeric types.

Cannot combine increments: must be separate statements.

Only statements: cannot assign x++ to a variable in Go.

Works in loops, functions, goroutines, as long as statement is valid.

*/
