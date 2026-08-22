package main

import "fmt"

type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return e.message
}

func getError() error {
	var err *MyError = nil
	return err
}

func main() {
	err := getError()

	fmt.Println(err)
	fmt.Println(err == nil)
}

/*


Value Receiver in Go
A receiver connects a method to a type.

When the receiver is written as:

func (d Dog) Speak() {
    ...
}
Dog is a value receiver.

1. Basic example
package main

import "fmt"

type Dog struct {
    Name string
}

func (d Dog) Speak() {
    fmt.Println(d.Name, "says Woof")
}

func main() {
    dog := Dog{Name: "Tommy"}

    dog.Speak()
}
Output:

Tommy says Woof
Here:

func (d Dog) Speak()
means d receives a copy of the Dog value.

Conceptually:

dog
┌─────────────┐
│ Name: Tommy │
└─────────────┘
      │
      │ copy
      ▼
d
┌─────────────┐
│ Name: Tommy │
└─────────────┘
2. Value receiver gets a copy
Consider:

package main

import "fmt"

type User struct {
    Name string
}

func (u User) ChangeName() {
    u.Name = "Bob"
}

func main() {
    user := User{Name: "Alice"}

    user.ChangeName()

    fmt.Println(user.Name)
}
Output:

Alice
Why?

Because:

func (u User) ChangeName()
receives a copy.

Original                    Copy

user                        u
┌─────────────┐            ┌─────────────┐
│ Name: Alice │  ─copy──►  │ Name: Alice │
└─────────────┘            └─────────────┘

                             u.Name = Bob

Original remains:

Name: Alice
3. Pointer receiver behaves differently
Now:

package main

import "fmt"

type User struct {
    Name string
}

func (u *User) ChangeName() {
    u.Name = "Bob"
}

func main() {
    user := User{Name: "Alice"}

    user.ChangeName()

    fmt.Println(user.Name)
}
Output:

Bob
Because u is a pointer to the original object.

user
┌─────────────┐
│ Name: Alice │
└─────────────┘
       ▲
       │
       │ pointer
       │
       u
Therefore:

u.Name = "Bob"
changes the original user.

4. Why can we call a value receiver using a pointer?
This is an important Go rule.

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof")
}

func main() {
    dog := Dog{}

    dog.Speak()

    p := &dog
    p.Speak()
}
Both work:

dog.Speak()
p.Speak()
Go automatically dereferences the pointer when necessary.

Conceptually:

p.Speak()
becomes approximately:

(*p).Speak()
because Speak() has a value receiver.

5. Can a value call a pointer receiver?
Yes, if the value is addressable.

type Dog struct{}

func (d *Dog) Speak() {
    fmt.Println("Woof")
}

func main() {
    dog := Dog{}

    dog.Speak()
}
Go automatically takes the address:

(&dog).Speak()
So this works.

6. But this doesn't always work
Consider:

Dog{}.Speak()
if Speak() has a pointer receiver:

func (d *Dog) Speak() {}
This doesn't work because Dog{} is not an addressable variable.

You can do:

dog := Dog{}
dog.Speak()
because dog is addressable.

7. Value receiver and interfaces
This is very important for interviews.

type Speaker interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof")
}
Now:

var s1 Speaker = Dog{}
var s2 Speaker = &Dog{}
Both are valid.

Why?

Because a value receiver method belongs to the method sets of:

Dog
and:

*Dog
Conceptually:

Dog
 └── Speak()

*Dog
 └── Speak()
Therefore both satisfy:

Speaker
8. Pointer receiver and interfaces
Now change the method:

func (d *Dog) Speak() {
    fmt.Println("Woof")
}
Then:

var s1 Speaker = Dog{}   // ❌
var s2 Speaker = &Dog{}  // ✅
Why?

The method set of:

Dog
doesn't contain pointer-receiver methods.

But:

*Dog
does.

So:

Dog
 └── no Speak()

*Dog
 └── Speak()
Therefore only:

&Dog{}
implements Speaker.

9. Method-set table
This is worth memorizing:

Receiver	T has method?	*T has method?
func (t T)	✅	✅
func (t *T)	❌	✅
Example:

func (d Dog) A() {}
func (d *Dog) B() {}
Method sets:

Dog:
    A()

*Dog:
    A()
    B()
This explains many interface questions.

10. When should you use a value receiver?
Use a value receiver when:

1. The type is small
For example:

type Point struct {
    X int
    Y int
}
func (p Point) Distance() float64 {
    ...
}
2. The method doesn't need to modify the object
func (u User) FullName() string {
    return u.FirstName + " " + u.LastName
}
3. The type behaves like a value
Examples:

Point
Coordinate
Money
Duration-like small structs
Configuration values
11. When should you use pointer receiver?
Use:

func (u *User) ...
when the method needs to modify the receiver:

func (u *User) SetName(name string) {
    u.Name = name
}
Or when copying the struct is expensive:

type LargeObject struct {
    Data [1000000]byte
}
Using:

func (x LargeObject) Process() {}
would copy a potentially huge value.

A pointer receiver avoids that copy:

func (x *LargeObject) Process() {}
12. Important subtlety: value receiver does not always mean "everything is copied"
Consider:

type User struct {
    Name string
    Tags []string
}
With:

func (u User) AddTag() {
    u.Tags[0] = "Go"
}
The User struct itself is copied, but Tags is a slice.

A slice internally contains:

pointer
length
capacity
So the copied slice can still refer to the same underlying array.

Therefore:

u.Tags[0] = "Go"
may modify the underlying data visible from the original object.

This is why "value receiver means nothing can affect the original" is not universally true.

13. Interview example
type Counter struct {
    Value int
}

func (c Counter) Increment() {
    c.Value++
}

func main() {
    c := Counter{}

    c.Increment()

    fmt.Println(c.Value)
}
Output:

0
Because:

c.Value = 0
    │
    │ copy
    ▼
receiver.Value = 0
    │
    └── increment → 1

original remains 0
Change it to:

func (c *Counter) Increment() {
    c.Value++
}
Now:

1
14. Value receiver vs pointer receiver
Feature	Value Receiver	Pointer Receiver
Syntax	(x T)	(x *T)
Gets copy	Yes	No
Can modify original struct	No	Yes
Avoids large struct copy	No	Yes
T satisfies interface	If method required	No
*T satisfies interface	Yes	Yes
Good for immutable/value-like behavior	Yes	Usually no
Most important interview statement
A value receiver operates on a copy of the receiver. Its method is in the method set of both T and *T, so both T and *T can satisfy an interface requiring that method.


if Speak() has a pointer receiver:

func (d *Dog) Speak() {}
This doesn't work because Dog{} is not an addressable variable.

You can do:

dog := Dog{}
dog.Speak()

because dog is addressable.
addressable means


What does "addressable" mean in Go?
In simple terms:

A value is addressable if Go can take its memory address using &.

For example:

dog := Dog{}
You can do:

fmt.Println(&dog)
Therefore, dog is addressable.

1. Simple example
type Dog struct {
    Name string
}

func (d *Dog) Speak() {
    fmt.Println(d.Name)
}

func main() {
    dog := Dog{Name: "Tommy"}

    dog.Speak()
}
You wrote:

dog.Speak()
but the method requires:

*d
Go effectively does:

(&dog).Speak()
Why can Go do that?

Because dog is addressable:

&dog
is valid.

2. Why is Dog{} different?
Consider:

Dog{}.Speak()
Suppose:

func (d *Dog) Speak() {
    fmt.Println("Woof")
}
Go would need to do:

(&Dog{}).Speak()
But Dog{} is a temporary value, not an addressable variable in this context.

So this doesn't work.

The important difference is:

dog := Dog{}
Here dog has a storage location:

dog
 │
 ▼
┌─────────────┐
│ Name: Tommy │
└─────────────┘
Therefore:

&dog
is possible.

But:

Dog{}
is just a value expression:

Dog{}
  │
  ▼
┌─────────────┐
│ Name: ""    │
└─────────────┘
It isn't addressable in this method-call situation.

3. Think of addressable as "has an identifiable memory location"
For example:

x := 10

fmt.Println(&x)
Works.

x is addressable.

You can also modify it:

x = 20
because x represents a specific storage location.

Compare:

10
You can't do:

&10 // ❌
because 10 is a literal value, not an addressable variable.

4. Examples
Addressable
x := 10

&x
✅

dog := Dog{}

&dog
✅

arr := [3]int{1, 2, 3}

&arr[0]
✅

An array element is addressable.

p := &Dog{}

&p
✅ p itself is addressable.

Not addressable
&10
❌

&(x + 1)
❌

&Dog{}
This is an important distinction: taking the address of a composite literal is allowed.

p := &Dog{}
✅

But:

Dog{}.Speak()
with a pointer receiver is not allowed because Go's automatic address-taking for method calls only applies when the value is addressable.

5. Why does Go automatically take the address?
Suppose:

type Dog struct{}

func (d *Dog) Speak() {
    fmt.Println("Woof")
}
You have:

dog := Dog{}

dog.Speak()
Go sees:

Speak requires *Dog
dog is Dog
dog is addressable
So Go allows:

dog.Speak()
as shorthand for:

(&dog).Speak()
This is called implicit address taking for method calls.

6. Compare these three
Case 1 — works
dog := Dog{}
dog.Speak()
Because:

dog → addressable
Case 2 — works
dog := Dog{}
(&dog).Speak()
Explicitly passing a pointer.

Case 3 — doesn't work
Dog{}.Speak()
Because Go cannot implicitly turn this into:

(&Dog{}).Speak()
for the purpose of method invocation.

However, this is valid:

(&Dog{}).Speak()
because you're explicitly creating a pointer to the composite literal.

7. Very important interview rule
Remember this:

If a method has a pointer receiver, a value can call that method only when the value is addressable. Go automatically takes its address.

So:

dog := Dog{}
dog.Speak()       // ✅
(&dog).Speak()    // ✅
(&Dog{}).Speak()  // ✅
Dog{}.Speak()     // ❌
assuming Speak() is:

func (d *Dog) Speak() {}
One-line definition
Addressable = a value that has a storage location whose address Go can obtain, typically with &.


Today 2:32 PM
give all trickyquestion of interface

Absolutely. For a senior Go interview, interfaces have many tricky areas beyond just type X interface{}. Below is a high-value set of tricky interface questions with answers, code, and expected output.

Go Interface — Tricky Interview Questions
1. What is an interface?
An interface defines a set of methods.

type Speaker interface {
    Speak()
}
A type satisfies an interface implicitly by implementing all required methods.

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof")
}
No explicit:

implements Speaker
is required.

2. The most important question: Is a nil interface really nil?
var x interface{}

fmt.Println(x == nil)
Output
true
Because:

x
┌───────────────┐
│ Type  = nil   │
│ Value = nil   │
└───────────────┘
3. Typed nil interface
Very common interview question:

var p *int = nil

var x interface{} = p

fmt.Println(x == nil)
Output
false
Why?

x
┌─────────────────┐
│ Type  = *int    │
│ Value = nil     │
└─────────────────┘
The interface has a dynamic type, so the interface itself isn't nil.

Remember
nil interface:

type  = nil
value = nil

typed nil:

type  = *int
value = nil
4. Why is this dangerous with error?
type MyError struct{}

func (e *MyError) Error() string {
    return "error"
}

func getError() error {
    var e *MyError = nil
    return e
}

func main() {
    err := getError()

    if err != nil {
        fmt.Println("Error!")
    }
}
Output:

Error!
Even though the pointer itself is nil.

Because:

error interface
┌─────────────────┐
│ Type  = *MyError│
│ Value = nil     │
└─────────────────┘
5. Value receiver vs pointer receiver
type Dog struct{}

func (d Dog) Speak() {}
What implements the interface?

type Speaker interface {
    Speak()
}
Both:

var s1 Speaker = Dog{}
var s2 Speaker = &Dog{}
work.

Why?

A value receiver belongs to the method sets of both:

Dog
*Dog
6. Pointer receiver — only pointer implements interface
func (d *Dog) Speak() {}
Now:

var s Speaker = Dog{}
❌ Compile error.

But:

var s Speaker = &Dog{}
✅ Works.

Method sets:

Dog:
    no Speak()

*Dog:
    Speak()
7. Tricky: Why does this work then?
You might ask:

dog := Dog{}

dog.Speak()
when:

func (d *Dog) Speak() {}
This works because dog is addressable.

Go effectively does:

(&dog).Speak()
But:

Dog{}.Speak()
doesn't work because Dog{} isn't addressable for implicit address-taking.

8. Interface equality
var a interface{} = 10
var b interface{} = 10

fmt.Println(a == b)
Output:

true
Both contain:

type  = int
value = 10
9. Different dynamic types
var a interface{} = int(10)
var b interface{} = int64(10)

fmt.Println(a == b)
Output:

false
Even though numerically both are 10.

a → int
b → int64
Dynamic types are different.

10. Interface containing a slice
This is a very tricky one.

var a interface{} = []int{1, 2}
var b interface{} = []int{1, 2}

fmt.Println(a == b)
This causes a runtime panic.

Why?

Interface equality requires the dynamic values to be comparable.

Slices aren't comparable.

You can compare:

a == nil
but you cannot compare two slices with ==.

For slice contents:

reflect.DeepEqual(a, b)
or:

slices.Equal(...)
11. Which types are comparable?
Generally:

Comparable
bool
numbers
string
pointer
channel
interface
array (if elements comparable)
struct (if all fields comparable)
Not comparable
slice
map
function
Therefore:

var x interface{} = []int{1, 2}

fmt.Println(x == x)
can panic because the dynamic value is a slice.

12. Interface containing map
var x interface{} = map[string]int{
    "a": 1,
}

fmt.Println(x == nil)
This is:

false
But:

fmt.Println(x == x)
causes a runtime panic because the dynamic value is a map, which is not comparable.

13. Interface containing function
var x interface{} = func() {}

fmt.Println(x == nil)
This is:

false
But comparing two interfaces containing functions causes a panic because functions aren't comparable except against nil.

14. Empty interface
Before generics, you often saw:

interface{}
Now you can write:

any
These are equivalent:

var x interface{}
var y any
any is simply an alias for:

interface{}
15. any can contain anything
var x any

x = 10
x = "hello"
x = true
x = []int{1, 2, 3}
x = map[string]int{"a": 1}
All are valid.

But remember:

An interface doesn't mean the value has no type. It means the interface can hold values of different dynamic types.

16. Type assertion
Suppose:

var x any = "Hello"
You can retrieve the underlying string:

s := x.(string)

fmt.Println(s)
Output:

Hello
This is called a type assertion.

17. Dangerous type assertion
var x any = "Hello"

n := x.(int)
This causes:

panic: interface conversion: interface {} is string, not int
Because the dynamic type is:

string
not:

int
18. Safe type assertion
Use the comma-ok form:

n, ok := x.(int)

if ok {
    fmt.Println(n)
} else {
    fmt.Println("Not an int")
}
Output:

Not an int
This is the preferred approach when the type isn't guaranteed.

19. Type switch
Another important interview topic:

func printType(x any) {
    switch v := x.(type) {
    case int:
        fmt.Println("int:", v)

    case string:
        fmt.Println("string:", v)

    case bool:
        fmt.Println("bool:", v)

    default:
        fmt.Println("unknown")
    }
}
Usage:

printType(10)
printType("hello")
printType(true)
20. Tricky type switch with nil
func check(x any) {
    switch x.(type) {
    case nil:
        fmt.Println("nil")

    default:
        fmt.Println("not nil")
    }
}
Then:

var x any

check(x)
Output:

nil
But:

var p *int = nil
var x any = p

check(x)
Output:

not nil
Because x has dynamic type *int.

21. Interface embedding
type Reader interface {
    Read()
}

type Writer interface {
    Write()
}

type ReadWriter interface {
    Reader
    Writer
}
A type implementing ReadWriter must implement both:

Read()
Write()
22. Interface can contain another interface?
An interface definition can embed another interface:

type Animal interface {
    Speaker
    Mover
}
But an interface value itself contains a dynamic concrete value, not another interface as its dynamic type in the ordinary assignment model.

This is a subtle distinction worth remembering:

interface definition → can embed interfaces

interface value → dynamic type is normally a concrete non-interface type
23. Interface assignment
type Speaker interface {
    Speak()
}

type Dog struct{}

func (Dog) Speak() {}

var d Dog

var s Speaker = d
The interface stores:

s
┌───────────────┐
│ Type  = Dog   │
│ Value = d     │
└───────────────┘
24. Copying an interface
var a any = 10
var b any = a

fmt.Println(a == b)
Output:

true
Both interface values contain the same dynamic value.

But remember that interface assignment doesn't necessarily deep-copy the underlying object.

25. Interface containing pointer
type User struct {
    Name string
}

u := &User{Name: "Alice"}

var x any = u
var y any = x

y.(*User).Name = "Bob"

fmt.Println(u.Name)
Output:

Bob
Why?

Both interfaces contain a pointer to the same User.

x ──────┐
        ▼
      User
        ▲
        │
y ──────┘
26. Interface method can have pointer receiver
type Speaker interface {
    Speak()
}

type Dog struct {
    Name string
}

func (d *Dog) Speak() {
    fmt.Println(d.Name)
}
This is valid:

var s Speaker = &Dog{Name: "Tommy"}
But:

var s Speaker = Dog{Name: "Tommy"}
doesn't compile.

27. Interface with unexported method
This is a powerful design technique.

type sealed interface {
    doSomething()
}
A package can use an unexported method to prevent external packages from implementing the interface.

For example:

type MyInterface interface {
    Do()
    private()
}
External packages can't implement private() because they cannot define the same unexported method belonging to your package.

This is sometimes used to create a closed/sealed interface pattern.

28. Interface segregation
Bad:

type Worker interface {
    Work()
    Eat()
    Sleep()
    Drive()
}
A type may only need:

Work()
Better:

type Worker interface {
    Work()
}

type Driver interface {
    Drive()
}
This follows the Interface Segregation Principle.

Go generally favors small interfaces.

A famous Go philosophy is:

Accept interfaces, return concrete types.

Not an absolute rule, but a useful design guideline.

29. Interface as function parameter
Instead of:

func Save(d Dog) {
}
you can define behavior:

type Saver interface {
    Save()
}

func Process(s Saver) {
    s.Save()
}
Now different types can participate:

type User struct{}
type Order struct{}

func (User) Save() {}
func (Order) Save() {}
Both can be passed to:

Process(User{})
Process(Order{})
30. Interface satisfaction check
You can explicitly verify that a type implements an interface:

var _ Speaker = (*Dog)(nil)
This doesn't create a meaningful runtime object.

It is a compile-time check.

If Dog stops implementing Speaker, the compiler reports an error.

31. Tricky: Does Dog implement the interface?
type Speaker interface {
    Speak()
}

type Dog struct{}

func (d *Dog) Speak() {}
Question:

var _ Speaker = Dog{}
❌ No.

But:

var _ Speaker = (*Dog)(nil)
✅ Yes.

Because Speak() has a pointer receiver.

32. Interface vs concrete type
Suppose:

type Dog struct {
    Name string
}

type Speaker interface {
    Speak()
}
This:

var d Dog
has static type:

Dog
This:

var s Speaker = Dog{}
has:

static type  = Speaker
dynamic type = Dog
dynamic value = Dog{}
This distinction is extremely important.

33. Static type vs dynamic type
Consider:

var x any = 100
Then:

Static type:
any

Dynamic type:
int

Dynamic value:
100
You can see the dynamic type using:

fmt.Printf("%T\n", x)
Output:

int
34. Tricky interface hierarchy
type Animal interface {
    Speak()
}

type Dog struct{}

func (Dog) Speak() {}
Now:

var a Animal = Dog{}
You can do:

d := a.(Dog)
because the dynamic type is Dog.

But:

d := a.(*Dog)
will panic because the dynamic type isn't *Dog.

35. Dog vs *Dog type assertion
var a any = Dog{}
Then:

a.(Dog)   // ✅
a.(*Dog)  // ❌
If:

var a any = &Dog{}
then:

a.(*Dog)  // ✅
a.(Dog)   // ❌
The dynamic type must match the assertion.

36. Tricky: Interface method with nil receiver
type Dog struct{}

func (d *Dog) Speak() {
    fmt.Println("Woof")
}

func main() {
    var d *Dog = nil

    d.Speak()
}
This can work because the method doesn't dereference d.

Now:

func (d *Dog) Speak() {
    fmt.Println(d.Name)
}
could panic.

So:

Calling a method on a nil pointer doesn't automatically mean panic. Dereferencing the nil pointer does.

37. Nil interface inside another interface
Be careful with:

var x any = nil
var y any = x

fmt.Println(y == nil)
Output:

true
Because x has no dynamic type/value.

But:

var p *Dog = nil

var x any = p
var y any = x

fmt.Println(y == nil)
Output:

false
The dynamic type survives:

*Dog
38. Interface holding nil map
var m map[string]int

var x any = m

fmt.Println(m == nil)
fmt.Println(x == nil)
Output:

true
false
Again:

map = nil
interface = not nil
39. Interface holding nil channel
var ch chan int

var x any = ch

fmt.Println(ch == nil)
fmt.Println(x == nil)
Output:

true
false
Same principle.

40. Interface holding nil function
var f func()

var x any = f

fmt.Println(f == nil)
fmt.Println(x == nil)
Output:

true
false
This is the same typed-nil phenomenon.

41. The ultimate nil table
Value	Interface assignment	x == nil
nil	var x any = nil	✅ true
(*Dog)(nil)	var x any = p	❌ false
([]int)(nil)	var x any = s	❌ false
(map[string]int)(nil)	var x any = m	❌ false
(chan int)(nil)	var x any = ch	❌ false
(func())(nil)	var x any = f	❌ false
42. Tricky: Can an interface contain another interface?
Suppose:

type A interface {
    Speak()
}

type B interface {
    Speak()
}

var a A = Dog{}
var b B = a
This is valid.

But the dynamic value is still the concrete Dog, not "interface A".

Conceptually:

a
↓
dynamic type = Dog

b
↓
dynamic type = Dog
The interface-to-interface assignment checks whether the dynamic concrete value satisfies the destination interface.

43. Empty interface and type switch
func check(v any) {
    switch v := v.(type) {
    case int:
        fmt.Println("integer", v)

    case string:
        fmt.Println("string", v)

    case nil:
        fmt.Println("nil")

    default:
        fmt.Printf("unknown %T\n", v)
    }
}
This is very common in real Go code.

44. Interface and generics are not the same
A common interview question:

Interface
Used for behavior:

type Reader interface {
    Read()
}
Generic
Used for type parameterization:

func Print[T any](v T) {
    fmt.Println(v)
}
Think:

Interface → "What can you do?"

Generic → "What type are you?"
Although Go's type constraints also use interfaces, they serve a different role from ordinary runtime interface values.

45. Interface as generic constraint
Modern Go:

type Number interface {
    int | int64 | float64
}
Then:

func Add[T Number](a, b T) T {
    return a + b
}
Here Number is being used as a type constraint, not as an ordinary interface value.

You cannot simply do:

var x Number
in the same sense as a normal runtime interface, because constraint interfaces can specify type sets that aren't ordinary runtime interface types.

46. interface{} vs any
Exactly equivalent:

interface{}
and:

any
So:

var x interface{}
and:

var x any
mean the same thing.

any is usually more readable in modern Go.

47. Interface with zero methods
type Empty interface{}
Any type implements it:

var x Empty

x = 10
x = "hello"
x = Dog{}
x = []int{1, 2}
That's why:

any
can hold any value.

48. Tricky: Does a nil pointer satisfy an interface?
Suppose:

type Speaker interface {
    Speak()
}

type Dog struct{}

func (*Dog) Speak() {}
This:

var d *Dog = nil

var s Speaker = d
is valid.

The interface is not nil:

fmt.Println(s == nil)
Output:

false
The interface contains:

dynamic type  = *Dog
dynamic value = nil
49. Most important interview puzzle
What does this print?

package main

import "fmt"

type Dog struct{}

func (d *Dog) Speak() {
    fmt.Println("Woof")
}

func main() {
    var d *Dog = nil
    var s interface{} = d

    fmt.Println(d == nil)
    fmt.Println(s == nil)
}
Answer
true
false
This single example tests:

pointer nil

interface nil

dynamic type

typed nil

method sets

50. Another classic puzzle
What happens?

package main

import "fmt"

func main() {
    var x interface{} = []int{1, 2}
    var y interface{} = []int{1, 2}

    fmt.Println(x == y)
}
Answer
Runtime panic.

Because the dynamic type is []int, and slices aren't comparable.

51. Another classic puzzle
var x interface{} = 10

switch x.(type) {
case int:
    fmt.Println("int")
case interface{}:
    fmt.Println("interface")
default:
    fmt.Println("other")
}
Output:

int
The more specific case is matched first.

52. Another classic puzzle
type A interface {
    Foo()
}

type B interface {
    Foo()
}

type C struct{}

func (C) Foo() {}

var a A = C{}
var b B = a

fmt.Printf("%T\n", b)
Output:

main.C
Not:

main.A
The dynamic concrete type remains C.

53. Interview cheat sheet
Memorize these rules:

1. Go interfaces are implicitly implemented.

2. Interface has a static type and dynamic type/value.

3. nil interface:
       type = nil
       value = nil

4. Typed nil:
       type = *T
       value = nil

5. Typed nil interface != nil.

6. Value receiver:
       T  implements interface
       *T also implements interface

7. Pointer receiver:
       T  does NOT implement interface
       *T implements interface

8. A value can call pointer-receiver methods
   when the value is addressable.

9. Type assertion:
       x.(T)

10. Safe assertion:
       v, ok := x.(T)

11. Type switch:
       switch v := x.(type)

12. any == interface{}.

13. Interface comparison can panic if
    dynamic value is non-comparable.

14. Slice/map/function cannot be compared
    with another value using ==.

15. Small interfaces are preferred in Go.

16. Interfaces describe behavior, not data.

17. Interface-to-interface assignment preserves
    the underlying concrete dynamic value/type.

18. nil pointer can be stored inside a non-nil interface.

19. Interface satisfaction is compile-time.

20. Generic constraints and runtime interfaces
    are related but not the same thing.
⭐ The 5 questions I would expect most often in a senior Go interview
Why does var err error = (*MyError)(nil) make err != nil?

Why does Dog{} fail for a pointer-receiver interface but dog := Dog{}; dog.Speak() work?

What is the difference between the method set of T and *T?

Why can comparing two interfaces containing slices/maps panic?

What are static type, dynamic type, and dynamic value of an interface?

These five cover a surprisingly large portion of the tricky interface questions interviewers use.

*/
