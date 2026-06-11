package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	min := 0
	max := 100

	num := r.Intn(max-min+1) + min
	fmt.Println(num)
}
