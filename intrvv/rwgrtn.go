package main

import "fmt"

func Reader(c <-chan int, b chan<- int) {
	for val := range c {
		b <- val
	}
}

func main() {

	r := make(chan int)
	w := make(chan int)

	go Reader(r, w)
	go func() {
		for val := range w {
			fmt.Println(val)
		}
	}()

	for i := 0; i < 6; i++ {
		r <- i*i + i + 4
	}

}
