package main

import "fmt"

func main() {
	//default variable are initialize to zero or default value

	var myvar1 int
	var myvar2 string
	var myvar3 float32
	fmt.Println("The value of Variable Myvar1 is ", myvar1)
	fmt.Println("The value of Variable Myvar2 is ", myvar2)
	fmt.Println("The value of Variable Myvar3 is ", myvar3)
	fmt.Println("*********************************************")
	//declaring multiple variable by removing type
	var p1, p2, p3, p4 = 23, "shailesh", 456.987, 9999999999.87777777

	fmt.Println("The value of variable p1 is ", p1)
	fmt.Printf("Type of variable p1 is %T\n", p1)

	fmt.Println("The value of variable p2 is ", p2)
	fmt.Printf("Type of variable p2 is %T\n", p2)

	fmt.Println("The value of variable p3 is ", p3)
	fmt.Printf("Type of variable p3 is %T\n", p3)

	fmt.Println("The value of variable p4 is ", p4)
	fmt.Printf("Type of variable p4 is %T\n", p4)
	fmt.Println("------------------------------------------")
	// by default it take float64 format to store float

	// declaring and initialize at time of declaration without var and type
	// variablename := expression

	myvars1 := 345
	myvars2 := 3456.987
	fmt.Println("The value of variable myvars1 is ", myvars1)
	fmt.Println("The Value of Variable myvars2 is ", myvars2)

	fmt.Println("&&&*******End of File ************")
}
