package main

import (
	"fmt"
)

func main() {
	a, b := 5, 9

	swap(&a, &b)

	fmt.Println(a, b)
}

func swap(e, f *int) {
	*e, *f = *f, *e
}
