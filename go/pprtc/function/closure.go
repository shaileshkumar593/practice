package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	c := counter()

	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3
}

/*

	How Closure Works Internally
		Normally:

				Local variables live on stack

				When function returns → stack frame destroyed

		But in closure:

				Go compiler detects variable is captured

				Moves it to heap

				So it survives after function returns

				This is called escape analysis
*/

/*

	Memory Visualization
Without closure:

counter()
  count → stack
return → stack destroyed
With closure:

counter()
  count → heap (because inner func uses it)

Returned function keeps reference to heap memory
*/
