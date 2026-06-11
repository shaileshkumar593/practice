package main

import (
	"fmt"
)

type Vehicle struct {
}

type School struct {
	strength int
}

type College struct {
	strength int
}
type Bank struct {
	strength int
}

func main() {

	var st Vehicle

	if (Vehicle{} == st) {

		fmt.Println("It is empty structure")

	} else {

		fmt.Println("It is not empty structure")

	}
	var sc School
	if (School{2000} == sc) {

		fmt.Println("It is empty structure")

	} else {

		fmt.Println("It is not empty structure")

	}
	var coll College
	if (College{} == coll) {

		fmt.Println("It is empty structure")

	} else {

		fmt.Println("It is not empty structure")

	}
	var bnk Bank
	if (Bank{} == bnk) {

		fmt.Println("It is empty structure")

	} else {

		fmt.Println("It is not empty structure")

	}
}
