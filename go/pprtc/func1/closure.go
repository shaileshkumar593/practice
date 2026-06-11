package main

import "fmt"

func Counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	c := Counter()

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

}

/*
	Internally (Conceptually)
		Go allocates count on the heap because it survives after function return.

		This is called:

		escape analysis

		variable escaping to heap

		Because:

		count
		is captured by closure.


		Closures Store State
This is why closures are useful for:

		counters

		middleware

		caching

		rate limiting

		retries

		generators

		maintaining private state

*/
