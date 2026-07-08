package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil receiver>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	describe(i)

	// Check before calling the method
	if i != nil {
		i.M()
	} else {
		fmt.Println("i is nil")
	}

	// Assign a concrete value
	i = &T{"Hello Go"}
	describe(i)
	i.M()

	// Interface contains a nil pointer
	var t *T
	i = t
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}