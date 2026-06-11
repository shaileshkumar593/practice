package main

import (
	"fmt"
)

func main() {
	var e struct{}
	var e2 struct{}
	fmt.Printf("%p\n", &e)  // 0x90b418
	fmt.Printf("%p\n", &e2) // 0x90b418
	fmt.Println(&e == &e2)  // true
}

// No matter how many empty structs you create, they all point to the same address.
/*
	Why Zero Memory and Same Address?
To understand why empty structs have zero size and share the same address, we need to delve into Go’s source code.

/go/src/runtime/malloc.go

// base address for all 0-byte allocations
var zerobase uintptr

func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer {
  // ...
  if size == 0 {
    return unsafe.Pointer(&zerobase)
  }
  // ...
According to this excerpt from malloc.go, when the size of the object to allocate is 0,
it returns a pointer to zerobase. zerobase is a base address used for allocating zero-byte objects and doesn’t
occupy any actual memory space.*/
