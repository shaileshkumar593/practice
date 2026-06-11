// GOLANG PROGRAM TO FIND GCD USING RECURSION
package main

// fmt package provides the function to print anything
import "fmt"

// function prototype
func main() {

	// declare the variables
	var n1 int
	var n2 int
	fmt.Println("Golang Program to find GCD using recursion")

	// initialize the variables
	n1 = 36
	n2 = 60

	// print the result using in-built function fmt.Println()
	fmt.Println("G.C.D OF", n1, n2, "is", hcf(n1, n2))
}

// create the function hcf()
func hcf(n1 int, n2 int) int {
	if n2 != 0 {
		return hcf(n2, n1%n2)
	} else {
		return n1
	}
}
