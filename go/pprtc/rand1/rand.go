package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	items := []string{"apple", "banana", "cherry", "mango"}

	for i := 0; i < 15; i++ {
		radomItem := items[rand.Intn(len(items))]
		fmt.Println(radomItem)
	}

}
