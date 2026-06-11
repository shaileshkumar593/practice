package main

import (
	"fmt"
)

func main() {
	switch num := 70; {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		// fallthrough  -- error
	}

}
