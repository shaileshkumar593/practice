package main

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	// buffer full
	ch <- 3
}

/*package main

func main() {
	ch := make(chan int, 2)

	<-ch
}

Receiver waiting for data.

But:

nobody sends data

main goroutine blocked

👉 DEADLOCK

*/
