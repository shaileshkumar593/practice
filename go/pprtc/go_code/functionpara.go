package main

import "fmt"

// creating array and paasing as value and refrence

func main() {

	var arr = [6]int{67, 34, 70, 35, 50, 22}
	var aar2 = [5]int{75, 84, 37, 47, 79}

	var res int = sum(arr, aar2)
	fmt.Println("sum of array elements is ", res)
	fmt.Println("New array upadated in called function is", aar2)
}

func sum(a int, a2 int) int {
	a2[0] = 45
	var tot int
	for i := 0; i < len(a); i++ {
		tot = tot + a[i]

	}

}
