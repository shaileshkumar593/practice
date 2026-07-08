package main

import (
	"fmt"
	"sync"
)

func producer(name string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for i := 1; i <= 3; i++ {
			out <- fmt.Sprintf("%s-%d", name, i)
		}
	}()

	return out
}

func fanIn(channels ...<-chan string) <-chan string {
	out := make(chan string)

	var wg sync.WaitGroup

	for _, ch := range channels {

		wg.Add(1)

		go func(c <-chan string) {
			defer wg.Done()

			for value := range c {
				out <- value
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	p1 := producer("A")
	p2 := producer("B")
	p3 := producer("C")

	for value := range fanIn(p1, p2, p3) {
		fmt.Println(value)
	}
}