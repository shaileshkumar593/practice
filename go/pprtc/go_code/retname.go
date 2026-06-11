package main

import "fmt"

func myfunc(p, q int) (rectangle int, square int) {
	rectangle = p * q
	square = q * q
	return
}

func main() {
	var area1, area2 = myfunc(9, 10)

	fmt.Println("Area of rectangle is ", area1)
	fmt.Println("Area of Square is ", area2)

}
