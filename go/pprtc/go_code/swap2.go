package main

import "fmt"



func main()  {
	var a int = 30
	var p int = 20

	a,p = p,a
	fmt.Printf("p = %d and a = %d", p,a)

	
}