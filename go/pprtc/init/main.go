package main

import (
	"fmt"
	"time"
)

var weekday string

func printhi() {
	fmt.Println("Hello how are u")
}

func init() {
	printhi()
	fmt.Println("hello is all fine")
	weekday = time.Now().Weekday().String()
}

func main() {
	fmt.Printf("Today is %s\n", weekday)
}
