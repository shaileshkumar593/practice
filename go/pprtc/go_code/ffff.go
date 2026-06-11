package main

import "fmt"

func main() {
	// variable declared and
	//initialized without expression
	// var variable_name type = expression

	var myvar1 = 20
	var myvar2 = "geekforgeek"
	var myvar3 = 34.80

	// fmt is for format specifier
	fmt.Printf("The value of myvar1 is  %d\n", myvar1)
	fmt.Printf("The value of myvar1 is %T\n ", myvar1)

	fmt.Printf("The value of myvar2 is  %s\n", myvar2)
	fmt.Printf("The value of myvar2 is %T\n ", myvar2)

	fmt.Printf("The value of myvar3 is  %f\n", myvar3)
	fmt.Printf("The value of myvar3 is %T\n ", myvar3)

	fmt.Println("The value of myvr1 is ", myvar1)
	fmt.Print("The value of myvar1 is ", myvar1)
	fmt.Println("End of program")

	//println print line without formating string with newline
	//printf need formating string to be specified
	//print print line without appeneding newline and formating strinhg
}
