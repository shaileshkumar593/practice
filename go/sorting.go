package main

import (
	"fmt"
)

func main() {
	c := []int{4, 8, 2, 9, 6, 22, 8}

	for i := 1; i < len(c); i++ {
		key := c[i]
		j := i - 1

		for j >= 0 && key >= c[j] {
			c[j+1] = c[j]
			j = j - 1
		}
		c[j+1] = key
	}

	fmt.Println(c)
}
