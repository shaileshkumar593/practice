Go: Why and Why Not – Common Rules


1️⃣ Variable Declarations
var x int
x = 5  // ✅ valid
y := 10 // ✅ valid
z := x + 5 // ✅ valid
Why: Go requires variables to be declared before use.

Why not: Using undeclared variables is invalid.

fmt.Println(a) // ❌ invalid, a not declared
2️⃣ Type Mismatch
var x int = 5
x = 5.5 // ❌ invalid
Why: Go is strongly typed, no implicit conversion.

Alternative: Explicit conversion:

x = int(5.5)
3️⃣ Assignments Between Types
var a int
var b int64
a = b // ❌ invalid
Why not: int and int64 are different types.

Alternative:

a = int(b)
4️⃣ Constants
const c = 5
c = 6 // ❌ invalid
Why not: Constants are immutable.

5️⃣ Using Nil
var p *int = nil // ✅ valid
*p = 5           // ❌ invalid, nil pointer dereference
Why not: You cannot dereference a nil pointer.

6️⃣ Slices vs Arrays
arr := [3]int{1,2,3}
arr[0] = 10      // ✅ valid
arr[3] = 5       // ❌ invalid, index out of range
7️⃣ Slice Append
s := []int{1,2}
s = append(s, 3) // ✅ valid
Why not: Cannot append to arrays directly ([3]int is fixed size).

8️⃣ Map Access
m := map[string]int{"a":1}
v := m["a"]       // ✅ valid
v := m["b"]       // ✅ returns zero value
m["b"] = 2        // ✅ valid
delete(m, "b")    // ✅ valid
Why not: Cannot index nil map for assignment:

var m map[string]int
m["a"] = 1 // ❌ panic
9️⃣ Function Calls
func add(a, b int) int { return a + b }
x := add(5, 6) // ✅ valid
x := add(5)    // ❌ invalid, wrong arg count
10️⃣ Variadic Functions
func sum(nums ...int) int { return nums[0] }
sum(1,2,3) // ✅ valid
sum()      // ✅ valid, nums is empty slice
11️⃣ Type Conversion
var a int = 10
var b float64 = float64(a) // ✅ valid
b = a // ❌ invalid
12️⃣ Using _ Blank Identifier
_, y := 1, 2 // ✅ ignores first value
Why: Avoids unused variable errors.

13️⃣ Multiple Return Values
func swap(a, b int) (int, int) { return b, a }
x, y := swap(1,2) // ✅ valid
x := swap(1,2)    // ❌ invalid, too many return values
14️⃣ Loops
for i := 0; i < 5; i++ {} // ✅ valid
for i := range slice {}    // ✅ valid
for i := 0; i < 5; i += 2 {} // ✅ valid
for i := 0; i < 5; i = i++ {} // ❌ invalid, ++ is statement
15️⃣ Control Flow
if x := f(); x > 5 {} // ✅ valid
if x > 5 {}           // ✅ valid
if x := f(); x > 5; x < 10 {} // ❌ invalid, Go does not support multiple conditions in if header
16️⃣ Switch / Case
switch x { case 1: } // ✅ valid
switch { case x>5: } // ✅ valid
Why not: Each case must be comparable if switching on value type.

17️⃣ Pointers
x := 5
p := &x // ✅ valid
*p = 10 // ✅ valid
p = nil // ✅ valid
Why not: Dereferencing nil → runtime panic.

18️⃣ Structs
type Point struct { X, Y int }
p := Point{1,2}
p.X = 10 // ✅ valid
p.Z = 5  // ❌ invalid, field Z does not exist
19️⃣ Interfaces
var i interface{} = 5
i = "hello" // ✅ valid, interface can hold any type
Why not: Cannot perform operations on i without type assertion:

i++ // ❌ invalid
20️⃣ Nil Checks
var m map[string]int
if m == nil {} // ✅ valid
21️⃣ Error Handling
err := doSomething()
if err != nil { fmt.Println(err) } // ✅ valid
Why not: Cannot ignore errors if assignment returns multiple values and unused.

22️⃣ Slices and Capacity
s := make([]int, 2, 5) // ✅ len 2, cap 5
s = append(s, 1,2,3)   // ✅ len 5
Why not: Cannot assign beyond capacity without append.

23️⃣ Constants in Expressions
const c = 5
x := c + 2 // ✅ valid
c++        // ❌ invalid, constant immutable
24️⃣ Go Does Not Allow Implicit Conversions
var i int
var f float64
i = f // ❌ invalid
i = int(f) // ✅ valid
25️⃣ Using Maps / Slices / Channels Without Init
var m map[string]int
m["a"] = 1 // ❌ panic
var s []int
s[0] = 1   // ❌ panic
Why: Must initialize maps/slices before assignment.

✅ Summary
Go is strongly typed, explicit, and safe:

Variables must be declared before use.

Type conversions must be explicit.

Constants are immutable.

++ / -- are statements, not expressions.

Nil values (pointers, slices, maps) cannot be dereferenced.

Loops, slices, maps, structs, interfaces have strict rules.