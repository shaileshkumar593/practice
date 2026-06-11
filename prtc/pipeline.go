package main

import "fmt"

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < len(nums); i++ {
			out <- nums[i]
		}
		close(out)
	}()

	return out
}

func Square(ch <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for val := range ch {
			out <- val * val
		}
		close(out)
	}()

	return out
}

func main() {
	ch := generator(3, 5, 3, 7, 2, 9)

	out := Square(ch)

	for val := range out {
		fmt.Println(val)
	}

}
