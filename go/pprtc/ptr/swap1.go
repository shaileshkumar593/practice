/*

	| Concept         | Syntax / Example      | Notes                             |
| --------------- | --------------------- | --------------------------------- |
| Declare pointer | `var p *int`          | Default nil                       |
| Address of      | `p = &x`              | Gets memory address               |
| Dereference     | `*p`                  | Access value at pointer           |
| Swap values     | `*i, *j = *j, *i`     | Changes original values           |
| `new`           | `p := new(int)`       | Allocates memory, returns pointer |
| Struct pointer  | `n := &Node{val: 10}` | Modifies original struct          |
| Double pointer  | `var pp **int = &p`   | Pointer to pointer                |
| Slice/Map       | `s := []int{1,2,3}`   | Slices & maps are reference types |
| Nil check       | `if p == nil`         | Avoid runtime panic               |

*/

package main

import "fmt"

func main() {
	a, b := 99, 44
	fmt.Println("Before swap:", a, b)

	// Anonymous function using pointers
	func(i, j *int) {
		*i, *j = *j, *i
	}(&a, &b)

	fmt.Println("After swap:", a, b)
}
