package main

import "fmt"

func FanIn(ch1 <-chan int, ch2 <-chan int, out chan<- int) {

	for val := range ch1 {
		out <- val
	}

	for val := range ch2 {
		out <- val
	}

	close(out)
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	out := make(chan int)

	go FanIn(ch1, ch2, out)

	go func() {
		for i := 0; i < 7; i++ {
			ch1 <- i
			ch2 <- i * i
		}

		close(ch1)
		close(ch2)
	}()

	for val := range out {
		fmt.Println(val)
	}
}

// why send in go routine
/*


We put the producer code inside a goroutine because otherwise main gets blocked before it can start consuming from out.

The producer is moved into a separate goroutine so that sending to ch1 and ch2 can happen concurrently with reading from out.
With unbuffered channels, a send operation blocks until another goroutine receives the value. If the producer runs in main,
it may block before main starts consuming from out, causing a deadlock. Running the producer in its own goroutine allows producer,
FanIn, and consumer to execute concurrently.


*/
