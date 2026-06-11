package main

import "fmt"

//  return channels because each function becomes an independent pipeline stage.

/*

	Why does for range need close?
		Because receiver must know:

		“No more values are coming.”

		Without close:

		receiver waits forever

		deadlock occurs

*/

func Generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		defer close(out)

		for n := range in {
			out <- n * n
		}
	}()

	return out
}

func doubler(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {
			out <- n * 2
		}
	}()

	return out
}

func main() {

	gen := Generator(1, 2, 3, 4)

	sq := square(gen)

	db := doubler(sq)

	for val := range db {
		fmt.Println(val)
	}
}
