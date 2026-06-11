package main

import "fmt"

func main() {
	var number int64
	var sum int64

	for i := 0; i < 5; i++ {
		fmt.Println("Enter number ")
		val, err := fmt.Scanf("%d", &number)
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += number
		fmt.Println("Values entered is ", val, number)
	}
	fmt.Println("Sum of values entered is ", sum)
}
