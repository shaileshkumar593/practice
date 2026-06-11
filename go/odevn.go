package main

import (
	"fmt"
	"sync"
)

var target = 14

func EvenC(odd, even, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	n := 0
	out <- n

	for {
		n++

		if n > target {
			close(odd)
			return
		}

		odd <- n

		next, ok := <-even
		if !ok {
			close(odd)
			return
		}

		n = next
	}
}

func OddC(odd, even, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		n, ok := <-odd
		if !ok {
			close(even)
			return
		}

		out <- n

		if n >= target {
			close(even)
			return
		}

		n++

		even <- n
	}
}

func main() {
	odd := make(chan int)
	even := make(chan int)
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go EvenC(odd, even, out, &wg)
	go OddC(odd, even, out, &wg)

	go func() {
		wg.Wait()
		close(out)
	}()

	for v := range out {
		fmt.Println(v)
	}
}
