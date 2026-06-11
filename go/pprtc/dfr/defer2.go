package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("On Exit from Main")
	fmt.Println("Entry in Main ")
	fmt.Println("Continue in main")
	fmt.Println("On Exitting from Main")
}
