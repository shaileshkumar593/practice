package main

import (
	"fmt"
)

func main() {
	s := "shailesh"

	for _, s := range s {
		fmt.Printf("String is  %s and type is  %T \n", s, s)
	}

}
