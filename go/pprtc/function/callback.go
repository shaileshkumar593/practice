package main

import (
	"fmt"
)

func process(a int, callback func(int)) {
	result := a * 2
	callback(result)
}

func main() {
	process(5, func(r int) {
		fmt.Println("Result:", r)
	})
}

/*

	 Used in:

		HTTP middleware

		Worker pools

		Kafka consumers

		Event processing

*/
