package main

import (
	"fmt"
)

func main() {
	arr := [5]int{10, 20, 30, 40, 50}

	s1 := arr[1:4]
	fmt.Println("s1:", s1)
	s2 := append(s1, 60)

	s1[0] = 99
	s3 := append(s1, 88, 96)

	//s1[0] = 99

	// Address of whole array
	/*fmt.Printf("Address of array        : %p\n", &arr)

	fmt.Println("\nAddress of each element:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Address of arr[%d]      : %p\n", i, &arr[i])
	}*/

	fmt.Println("arr:", arr)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)

}

/*

s1: [20 30 40]
arr: [10 99 30 40 60]
s1: [99 30 40]
s2: [99 30 40 60]
s3: [99 30 40 88 96]
*/
