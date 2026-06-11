package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println("read single integer")
	p, err := fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Println("error ", err)
	}
	fmt.Println(p)

	var y = make([]int, 10)
	fmt.Println("read Array of int ")
	for i := 0; i < 7; i++ {
		k, err := fmt.Scanf("%d", &y[i])
		if err != nil {
			fmt.Println("------", err)
		}
		fmt.Println("count  ", k)
	}
	fmt.Println("*****************Scan******************")
	var x1 int
	fmt.Println("read single integer")
	p1, err := fmt.Scan(&x1)
	if err != nil {
		fmt.Println("error ", err)
	}
	fmt.Println(p1)

	var y1 = make([]int, 10)
	fmt.Println("read Array of int ")
	for i := 0; i < 7; i++ {
		k2, err := fmt.Scan(&y1[i])
		if err != nil {
			fmt.Println("------", err)
		}
		fmt.Println("count  ", k2)
	}

	fmt.Println("*****************Scanln******************")
	var x3 int
	fmt.Println("read single integer")
	p3, err := fmt.Scanln(&x3)
	if err != nil {
		fmt.Println("error ", err)
	}
	fmt.Println(p3)

	var y3 = make([]int, 10)
	fmt.Println("read Array of int ")
	for i := 0; i < 7; i++ {
		k, err := fmt.Scanf("%d", &y3[i])
		if err != nil {
			fmt.Println("------", err)
		}
		fmt.Println("count  ", k)
	}
}
