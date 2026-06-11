package main

import "fmt"

func generator(nums ...int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func multiplyBy2(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {
			out <- n * 2
		}
	}()

	return out
}

func add1(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {
			out <- n + 1
		}
	}()

	return out
}

func main() {

	ch1 := generator(1, 2, 3)

	ch2 := multiplyBy2(ch1)

	ch3 := add1(ch2)

	for result := range ch3 {
		fmt.Println(result)
	}
}
