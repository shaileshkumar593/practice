package main

import (
	"fmt"
)

func myfunc(ch chan string) {
	fmt.Println(<-ch)
	ch <- "Hell World"
}

func main() {
	ch := make(chan string)

	//ch <- "World Is Awsome to Live" /// deadlock condition sent on channel but no goroutine
	//all goroutines are asleep - deadlock!
	// main is block after sending but no receiver is there therefore it causes deadlock...
	go myfunc(ch) // first call receiver then send data through channel.
	ch <- "World Is Awsome to Live"
	val, ok := <-ch

	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("exiting from main")
	}
}
