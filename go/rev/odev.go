package main

import (
	"fmt"
)

var limit int = 15

func Evencnt(od, ev, out chan int) {

	for {
		val, ok := <-ev

		if !ok {
			close(od)
			return
		}

		out <- val

		val = val + 1

		if val > limit {
			close(od)
			return
		}
		od <- val
	}
}

func OddCnt(od, ev, out chan int) {

	for {
		val, ok := <-od

		if !ok {
			close(ev)
			return
		}

		out <- val

		val = val + 1

		if val > 14 {
			close(ev)
			return
		}
		ev <- val
	}
}

func main() {
	ev := make(chan int)
	od := make(chan int)
	out := make(chan int)

	go Evencnt(od, ev, out)
	go OddCnt(od, ev, out)

	ev <- 0

	for val := range out {
		fmt.Println(val)
		if val >= limit {
			close(out)
		}
	}
}
