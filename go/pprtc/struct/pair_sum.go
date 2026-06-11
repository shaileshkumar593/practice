package main

import "fmt"

func pairSum(nums []int, target int) bool {
	seen := make(map[int]struct{}) // set using empty struct

	for _, num := range nums {
		complement := target - num

		// check if complement exists
		if _, exists := seen[complement]; exists {
			return true
		}

		// store current number
		seen[num] = struct{}{}
	}

	return false
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(pairSum(nums, target)) // true
}

/*

	Why Do We Need It in Map?
When using a map as a set:

seen := make(map[int]struct{})
The value type is struct{}.

So when inserting:

seen[num] = struct{}{}
We must assign a value of type struct{}.
*/
/*

Why Not Just Write?
seen[num] = struct{}
Because:

struct{}
is a type, not a value.

Go requires a value.



struct{}
Has no fields → no storage needed.

The compiler optimizes it to zero-size type.


Important: Zero-Size Doesn’t Mean No Identity
Even though size is zero:

a := struct{}{}
b := struct{}{}
They are distinct values.

Because it:

Takes no memory

Is lightweight

Is commonly used as a signal or marker

*/
