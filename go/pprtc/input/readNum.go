package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("input number:")
	var number int64
	var Arry []int64
	var sum1 int64
	for i := 0; i < 5; i++ {
		_, err := fmt.Scan(&number)
		if err != nil {
			log.Fatal(err)
		}
		Arry = append(Arry, number)
		sum1 += number

	}
	fmt.Println(sum1)
	var sum int64
	for i := 0; i < len(Arry); i++ {
		fmt.Println("%d ---->%d", i, Arry[i])
		sum = sum + Arry[i]
	}
	fmt.Println("sum of %d element is %d ", len(Arry), sum)
}
