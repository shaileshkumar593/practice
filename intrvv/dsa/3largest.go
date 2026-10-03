package main

import "fmt"

func firstSecondThirdHighest(nums []int) (*int, *int, *int) {
	var first *int
	var second *int
	var third *int

	for _, num := range nums {

		// Skip duplicates
		if (first != nil && num == *first) ||
			(second != nil && num == *second) ||
			(third != nil && num == *third) {
			continue
		}

		// New highest
		if first == nil || num > *first {
			third = second
			second = first

			value := num
			first = &value

		} else if second == nil || num > *second {
			// New second highest
			third = second

			value := num
			second = &value

		} else if third == nil || num > *third {
			// New third highest
			value := num
			third = &value
		}
	}

	return first, second, third
}

func printValue(value *int) {
	if value == nil {
		fmt.Print("None")
	} else {
		fmt.Print(*value)
	}
}

func main() {
	testCases := [][]int{
		{10, 5, 20, 20, 8, 15, 3},
		{1, 2, 3},
		{10, 10, 10},
		{5, 4},
		{100},
		{},
		{-10, -5, -20, -1},
		{5, 5, 4, 4, 3, 3},
		{0, -1, -2, -3},
	}

	for _, nums := range testCases {
		first, second, third :=
			firstSecondThirdHighest(nums)

		fmt.Printf("nums=%v\n", nums)

		fmt.Print("1st=")
		printValue(first)

		fmt.Print(", 2nd=")
		printValue(second)

		fmt.Print(", 3rd=")
		printValue(third)

		fmt.Println()
		fmt.Println()
	}
}
