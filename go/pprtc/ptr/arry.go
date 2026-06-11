package main

import (
	"fmt"
)

// a[x] is shorthand for (*a)[x]. So (*arr)[0] in the above program can be replaced by arr[0]
func modify(arr *[3]int) {
	(*arr)[0] = 90
}

func main() {
	a := [3]int{89, 90, 91}
	modify(&a)
	fmt.Println(a)
}
