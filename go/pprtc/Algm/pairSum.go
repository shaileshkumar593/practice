package main

import "fmt"

/*
	store each element with index in map
	if sum -val == key of map then print element exist
	print sum -a[i], a[i]

*/

func main() {
	//a := []int{1, -2, 1, 0, 5}
	//a := []int{5, 0, 4, 3, 6, 2, 0, 3}----5
	a := []int{5, 0, 6, 3, 9, 8, 2, -4, -3}

	m := make(map[int]int, len(a))
	targetSum := 5
	count := 0
	for i := 0; i < len(a); i++ {
		if _, found := m[targetSum-a[i]]; found {
			count = count + 1
			fmt.Println(targetSum-a[i], "----->", a[i])
		} else {
			m[a[i]] = i

		}

	}
	if count > 0 {
		fmt.Println("Pair exist")
	} else {
		fmt.Println("Pair not exist")
	}

}
