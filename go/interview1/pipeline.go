package main

import "fmt"

// output of one is input to other
func generator(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func square(in chan int, out chan int) {
	for val := range in {
		out <- val * val
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go generator(ch1)
	go square(ch1, ch2)

	for result := range ch2 {
		fmt.Println(result)
	}
}
