package main

import "fmt"

type MyInt = int

func main() {
	var a MyInt = 100
	var b int = 200

	a = b // ✅ allowed
	b = a // ✅ allowed

	fmt.Println(a, b)
}

/*

"A type alias uses type A = B and makes A exactly another name for B,
so it doesn't create a new type or provide additional type safety.
A defined type uses type A B, creates a distinct type with B as its
underlying type, and allows methods and domain-specific type safety.
Aliases are particularly useful for backward compatibility and large-scale package refactoring."



var a A
var b B

a = b // error compile time

a = int(b)  // no error



type MyInt = int

func (x MyInt) Double() int {
	return int(x) * 2
}
This is invalid because MyInt is effectively int.

Can you add methods to an alias?

Answer:

It depends on what the alias denotes, but an alias does not create a new method-owning type.
You cannot use an alias to attach methods to an existing non-local type such as int or string.


but this works
type MyInt int

func (x MyInt) Double() int {
	return int(x) * 2
}

*/

/*
"In Go, type A = B creates a type alias, meaning A and B are exactly the same type.
It doesn't introduce a new type and is commonly used for API compatibility and package refactoring.
In contrast, type A B creates a new defined type whose underlying type is B. The new type provides
 compile-time type safety and can have its own methods. For example, I'd use type UserID int64 and type
  OrderID int64 in a backend domain model so the compiler prevents accidentally passing an OrderID where
   a UserID is expected. I'd use an alias such as type User = domain.User when migrating or
    reorganizing packages while maintaining backward compatibility."*/
