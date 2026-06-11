package main

import "fmt"

func main() {

	ch := make(chan string)

	// only  goroutine is allowed to talk not function
	go func() {
		ch <- "hello"
		ch <- "How"
		ch <- "are"
		ch <- "you"

	}()

	for val := range ch {
		fmt.Println(val)
	}

	/*
		keeps receiving from channel

		automatically stops when channel is closed

		If channel is NOT closed → infinite wait (deadlock)
	*/

}

/*
	PS C:\Users\shailesh.kumar\goconcurrency\channel> go run .\iterate_on_channel.go
		fatal error: all goroutines are asleep - deadlock!

*/
