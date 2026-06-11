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

	// Address of whole array
	/*fmt.Printf("Address of array        : %p\n", &arr)

	fmt.Println("\nAddress of each element:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Address of arr[%d]      : %p\n", i, &arr[i])
	}*/

	fmt.Println("arr:", arr)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

}

/*
	append within capacity → shares array

    append beyond capacity → new array
*/

/*

What happens when you slice an array?

Correct answer:

A slice header is created that points to the same underlying array. No data is copied.
*/

/*
Internally Slice Looks Like This
A slice is just:

type slice struct {
    ptr *T   // pointer to array
    len int
    cap int
}
So when you write:

s := arr[1:4]
Go creates something like:

ptr → &arr[1]
len = 3
cap = 4   (from index 1 to end of array)
📊 Memory Diagram
Initial array:

Index:  0   1   2   3   4
arr:   [10, 20, 30, 40, 50]
After slicing:

s := arr[1:4]

s:
ptr → arr[1]
len = 3
cap = 4
Visual:

arr:  [10, 20, 30, 40, 50]
           ↑   ↑   ↑
           s spans here
🔥 Important Consequence
Since slice shares same array:

s[0] = 99
This modifies:

arr[1]
Array becomes:

[10, 99, 30, 40, 50]
🚨 Very Important Rule
Slice is just a window over the array.

It does NOT copy data.

📌 Length vs Capacity
If:

s := arr[1:4]
Then:

len = 3  (4 - 1)
cap = 4  (5 - 1)
Because capacity is:

from start index to end of array
🔥 What Happens If You Append?
Case 1: Append within capacity
s2 := append(s, 60)
Since cap = 4 and len = 3 → space exists
So append modifies original array.

Case 2: Append exceeds capacity
s3 := append(s, 60, 70)
Now capacity exceeded.

Go creates NEW array.

Then:

s3 no longer shares arr
🧠 Important Difference
Operation	Result
s := arr[1:4]	Shares array
append() within cap	Still shares
append() beyond cap	New array created
🧨 Common Production Bug
func modify(s []int) {
    s[0] = 999
}
If slice came from array → original array changes.

Many developers don’t realize this.

🔥 If You Want Independent Copy
Do this:

sCopy := append([]int(nil), s...)
OR

sCopy := make([]int, len(s))
copy(sCopy, s)
Now:

sCopy has its own array

*/
