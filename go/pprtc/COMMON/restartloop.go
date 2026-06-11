package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}

	var count int
loop:
	for _, v := range slice {
		fmt.Println(v)

		count++
		if count == 2 {
			goto loop
		}
	}
}

/*
In the above snippet, the for..range loop is given a label
(loop, but it can be any name you want). Inside the loop,
the goto keyword is used to break out of the loop once a certain condition
is achieved. This moves control back to line 10 which effectively starts the
loop from the beginning.*/
