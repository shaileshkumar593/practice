package main

import "fmt"

func main() {
	var i int

	for i = 1; i < 10; i++ {
		fmt.Println("index no %d", i)
		fmt.Printf("index no2 %d\n", i)
		fmt.Print("index no3 ", i)
	}
}
