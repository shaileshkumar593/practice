package main

import "fmt"

func main() {
	set := make(map[int]struct{})

	set[10] = struct{}{}
	set[20] = struct{}{}

	if _, exists := set[10]; exists {
		fmt.Println("10 exists in set")
	}
}

/*
Used in:

Duplicate removal

Graph visited nodes

2Sum / 3Sum problems

Membership checks
*/
