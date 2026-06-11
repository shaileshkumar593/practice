package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch) reading from empty channel cause deadlock

	fmt.Println(len(ch), cap(ch))
} // working fine
