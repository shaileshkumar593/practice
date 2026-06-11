package main

import (
	"fmt"
)

func EvenCnt(ev, od chan int, out chan int) {

	for {
		val, ok := <-ev

		if !ok {
			close(od)
			return
		}

		out <- val

		val = val + 1

		if val > 15 {
			close(od)
			return
		}

		od <- val
	}
}

func OddCnt(ev, od chan int, out chan int) {

	for {
		val, ok := <-od

		if !ok {
			close(ev)
			//close(out) // close out once work is finished
			return
		}

		out <- val

		val = val + 1

		if val > 14 {
			close(ev)
			//close(out) // close out once work is finished
			return
		}

		ev <- val
	}
}

func main() {

	od := make(chan int)
	ev := make(chan int)
	out := make(chan int)

	l := make([]int, 0)

	go EvenCnt(ev, od, out)
	go OddCnt(ev, od, out)

	// Start sequence
	ev <- 0

	// Read values from out
	for val := range out {
		l = append(l, val)
		if val >= 15 {
			close(out)
		}
	}

	fmt.Println(l)
}
