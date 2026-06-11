package main

import "fmt"

func main() {
	//declare array
	a := [...]string{"shailesh", "priyanka", "Anish", "Gautam", "Golu"}
	fmt.Println(len(a))

	var c []string = a[1:4]

	fmt.Println("slice after c relation :", c)
	c[0] = "kanti"
	fmt.Println("slice after change is ", c)

	const h int = 7
	fmt.Println(h)

	var g []string = a[:]
	fmt.Println("g slice is ", g)

	/*	var t []string = a[:9]
		fmt.Println("t slice is ", t)

		Array out of bound exception

	*/
}
