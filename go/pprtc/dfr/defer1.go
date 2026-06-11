package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Hello World") // defer never execute because of exit
	fmt.Println("World is Jannat To Survive")
	//
	os.Exit(0)
}
