package main

import "fmt"

func Reader1(c <-chan int, b chan<- int) {
	for val := range c {
		b <- val
	}
	close(b)

}

func main() {

	r := make(chan int)
	w := make(chan int)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for val := range w {
			fmt.Println(val)
		}
	}()
	go Reader1(r, w)
	for i := 0; i < 6; i++ {
		r <- i*i + i + 4
	}
	close(r)
	<-done
}
