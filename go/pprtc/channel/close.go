package main

import "fmt"

func Myfunc(mychnl chan string) {
	fmt.Println(" MyFunc go routine is started ")
	for v := 0; v < 4; v++ {
		mychnl <- "GeeksforGeeks"
	}
	close(mychnl)
}

func main() {

	// Creating a channel
	c := make(chan string)

	// calling Goroutine
	go Myfunc(c)

	// When the value of ok is
	// set to true means the
	// channel is open and it
	// can send or receive data
	// When the value of ok is set to
	// false means the channel is closed
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}
