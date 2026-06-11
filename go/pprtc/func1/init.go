package main

import "fmt"

func init() {
	fmt.Println("Init executed")
}

func main() {
	fmt.Println("Main executed")
}

//Automatically runs before main().
