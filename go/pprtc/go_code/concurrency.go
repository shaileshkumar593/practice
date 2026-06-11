package main

import (
	"fmt"
	"reflect"
	"time"
)

func display() {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("In display")
	}
}

func main() {
	i := 78
	fmt.Println(reflect.TypeOf(i))
	rune := 98
	fmt.Println(rune)
	go display()

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("In the main")

	}
}
