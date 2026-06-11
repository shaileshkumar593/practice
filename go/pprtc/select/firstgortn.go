package main

import "fmt"

func PrintHello() {
	fmt.Println("In PrintHello go routine s")
	fmt.Println("Hello How are you")
}

func main() {

	go PrintHello()
	fmt.Println("In Main go routine ")

}
