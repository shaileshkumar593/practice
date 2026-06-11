package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //  with or without seed every run give unique value

	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(100) + 1)

	fmt.Println(rand.Float64())      //  0 <= f < 1.0
	fmt.Println(rand.Float64() * 10) // 0.0 to 10.0

	arr := []int{10, 20, 30, 40, 50}
	fmt.Println(arr[rand.Intn(len(arr))])

	const charSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	length := 12
	b := make([]byte, length)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}

	fmt.Println(string(b))

}
