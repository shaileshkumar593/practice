package main

import (
	"fmt"
	"time"
)

func producer(name string, nums []int, delay time.Duration) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, n := range nums {
			time.Sleep(delay)
			fmt.Println(name, "produced:", n)
			ch <- n
		}
	}()

	return ch
}

func fanIn(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				out <- v

			case v, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				out <- v
			}
		}
	}()

	return out
}

func main() {
	ch1 := producer("P1", []int{1, 2, 3}, 500*time.Millisecond)
	ch2 := producer("P2", []int{10, 20, 30}, 800*time.Millisecond)

	out := fanIn(ch1, ch2)

	for v := range out {
		fmt.Println("received:", v)
	}

	fmt.Println("done")
}
