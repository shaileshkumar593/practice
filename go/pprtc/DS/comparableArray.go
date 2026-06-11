package main

import "fmt"

func main() {
	a := [2]string{"strawberry", "raspberry"}
	b := [2]string{"strawberry", "raspberry"}
	fmt.Printf("slices equal: %t\n", a == b)
}
