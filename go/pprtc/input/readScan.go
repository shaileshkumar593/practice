package main

import "fmt"

func main() {
	var number int64
	var ele int64
	var sum int64
	fmt.Println("Enter number ")
	v, err := fmt.Scanf("%d", &ele)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v, ele)
	for i := 0; i < 5; i++ {
		fmt.Println("Enter number ")
		val, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += number
		fmt.Println("Values entered is ", val, number)
	}
	fmt.Println("Sum of values entered is ", sum)
}
