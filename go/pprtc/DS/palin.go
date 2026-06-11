package main

import "fmt"

func main() {

	var palNum, temp, reverse, palMin, palMax, remainder int

	fmt.Print("Enter the Minimum and Maximum limit of Palindrome = ")
	fmt.Scanln(&palMin, &palMax)

	fmt.Print("Palindrome Numbers between ", palMin, " and ", palMax, " are : ")
	for palNum = palMin; palNum <= palMax; palNum++ {
		reverse = 0
		for temp = palNum; temp > 0; temp = temp / 10 {
			remainder = temp % 10
			reverse = reverse*10 + remainder
		}
		if palNum == reverse {
			fmt.Print(palNum, "\t")
		}
	}
}
