/*
	If you need high-performance string manipulation frequently (e.g., building new strings in loops), prefer strings.Builder:
		1. strings.Builder is designed for efficient string concatenation.

		2. It minimizes memory allocations by managing an internal growing buffer.

		3. Unlike simple + string concatenation, this avoids creating multiple temporary strings.


		4. Use strings.Builder when you are building large or dynamic strings in loops.

		5. It’s not thread-safe, so don’t use it across multiple goroutines without synchronization.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder // Efficient string builder (no repeated allocations)

	b.WriteString("hello")
	b.WriteByte('!')

	fmt.Println(b.String()) // Output: hello!
}
