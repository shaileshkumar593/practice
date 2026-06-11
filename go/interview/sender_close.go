package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

/*
	close(ch)
	It means:

	 "No more values will be sent."

	It does NOT mean:

	Delete buffered values

	Immediately stop receivers

	Remove already-sent data

*/
/*
What Actually Happens Internally
For your code:

go func() {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}()
Step-by-step:

ch <- 1

Receiver gets 1

ch <- 2

Receiver gets 2

ch <- 3

Receiver gets 3

close(ch)

At the moment of close(ch):

All values already sent are still readable

Channel just marks an internal flag: closed = true

*/
