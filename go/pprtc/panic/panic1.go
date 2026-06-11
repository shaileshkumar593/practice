package main

import (
	"fmt"
)

func slicePanic() {
	n := []int{5, 7, 4}
	fmt.Println(n[4])
	fmt.Println("normally returned from a")
}
func main() {
	slicePanic()
	fmt.Println("normally returned from main")
}
