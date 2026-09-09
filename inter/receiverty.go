package main

import "fmt"

type MyInt1 int

func (n MyInt1) PrintFunc1() {
	fmt.Println("Value:", n)
}

func main() {
	n := MyInt1(100)
	n.PrintFunc1()
}
