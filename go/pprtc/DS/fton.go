package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the height in centimeters: ")
	fmt.Scanf("%d", &n)
	inches := 0.394 * float32(n)
	feet := 0.0328 * float32(n)
	fmt.Println("The length in inches:", inches)
	fmt.Println("The length in feet:", feet)
}
