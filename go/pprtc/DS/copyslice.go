package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{}
	slice2 = append(slice2, slice1...)

	fmt.Println(slice1) // => [1 2 3 4 5 6]
	fmt.Println(slice2) // => [1 2 3 4 5 6]

	//Modifying slice2 doesn't affect slice1
	slice2[0] = 100
	fmt.Println(slice1) // => [1 2 3 4 5 6]
	fmt.Println(slice2) // => [100 2 3 4 5 6]

	// copying a range of items
	slice3 := []int{}
	slice3 = append(slice3, slice1[3:5]...)
	fmt.Println(slice3) // => [4 5]

	//Again slice3 and slice1 doesn't share underlying array

	slice3[0] = -10
	fmt.Println(slice1) // => [1 2 3 4 5 6]
	fmt.Println(slice3) // => [-10 5]

}
