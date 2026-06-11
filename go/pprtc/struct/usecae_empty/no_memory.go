package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var e struct{}
	fmt.Println(unsafe.Sizeof(e)) // Output: 0
}

/*

1. No fields
2. Size = 0 bytes
3. Allocates no memory
*/
