package main

import (
	"fmt"
)

func number() int {
	num := 60 * 5
	return num
}

func main() {

	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)

	case num < 300:
		fmt.Printf("%d is lesser than 300", num)

	case num < 400:
		fmt.Printf("%d is lesser than 400", num)

	}

}
