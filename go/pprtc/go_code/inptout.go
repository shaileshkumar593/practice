package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var age int
	fmt.Println("Enter Age ")
	_, err := fmt.Scanf("%d", &age)

	fmt.Println("Error during reading age is ", err)
	reader := bufio.NewReader(os.Stdin)
	var name string
	fmt.Println("What is your name ")
	name, _ = reader.ReadString('\n')
	fmt.Println("Your name is ", name, " and your age is ", age)

}
