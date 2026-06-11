package main

import "fmt"

func main() {
	var scth int
	fmt.Println("choice for switch is 1,2,3")
	fmt.Print(" Enter sitch option ")
	fmt.Scanf("%d", &scth)

	switch scth {
	case 1:
		fmt.Println("The value of switch case is ", scth)

	case 2:
		fmt.Println("The value of switch case is ", scth)

	case 3:
		fmt.Println("The value of switch case is ", scth)

	default:
		fmt.Println("The value of default switch case is ", scth)

	}
}
