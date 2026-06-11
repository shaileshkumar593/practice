package main

import (
	"fmt"
	"math/rand"
	"time"
)

//"time"

func main() {
	fmt.Print(rand.Intn(100)) //will produce random integer between 0 to 100
	fmt.Println()

	fmt.Print(rand.Float64()) // will produce random number between 0 to 1
	fmt.Println()

	rand.Seed(time.Now().Unix()) // seeding do that random number will produced
	myrand := random(1, 20)

	fmt.Println(myrand)

}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
