/*
You are given an array of 0s and 1s in random order.
Segregate 0s on left side and 1s on right side of the array [Basically you have to sort the array]*/

package main

import "fmt"

// Segregate0And1 segregates 0s and 1s in-place in a slice
func Segregate0And1(arr []int) {
	lo, hi := 0, len(arr)-1

	for lo < hi {
		if arr[lo] == 1 {
			// Swap 1 at lo with 0 at hi
			if arr[hi] == 0 {
				arr[lo], arr[hi] = arr[hi], arr[lo]
			}
			hi-- // move hi pointer left
		} else {
			lo++ // move lo pointer right
		}
	}
}

func main() {
	arr := []int{0, 1, 0, 1, 1, 1}
	Segregate0And1(arr)
	fmt.Println("Array after segregation is:", arr)
}

/*
	Time Complexity: O(n)
	O(n) — each element is checked at most once.

	Space Complexity: O(1)
	O(1) — in-place swapping, no extra slices or arrays.
*/
