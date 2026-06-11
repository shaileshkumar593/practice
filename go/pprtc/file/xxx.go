package main

import "fmt"

/*
when new elements are appended to the slice, a new array is created.
The elements of the existing array are copied to this new array and
a new slice reference for this new array is returned. The capacity
of the new slice is now twice that of the old slice. Pretty cool
right :). The following program will make things clear.*/

func surprise(a []int) {
	a = append(a, 5)
	for i := range a {
		a[i] = 5
	}
	fmt.Println(a)
}

// Quiz #2

func main() {
	a := []int{1, 2, 3, 4}
	surprise(a)
	fmt.Println(a)
}

// a = [5,5,5,5,5]
//a = [1,2,3.4]
