package main

import "fmt"

func generate(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, n := range nums {
			out <- n
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

func double(in <-chan int) <-chan int {

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

	stage1 := generate(1,2,3,4)

	stage2 := square(stage1)

	stage3 := double(stage2)

	for value := range stage3 {
		fmt.Println(value)
	}
}