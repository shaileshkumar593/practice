package main

import (
	"fmt"
	"sync"
)

func Even(wg *sync.WaitGroup, ec chan int) {
	defer wg.Done()

	val := 0
	fmt.Println(val)

	for {
		val = val + 1
		ec <- val

		val, ok := <-ec
		if !ok {
			return
		}

		if val == 10 {
			fmt.Println(val)
			close(ec)
			return
		}
	}
}

func Odd(wg *sync.WaitGroup, oc chan int) {
	defer wg.Done()

	for {
		val, ok := <-oc
		if !ok {
			return
		}

		fmt.Println(val)

		if val == 9 {
			oc <- val + 1
			return
		}

		oc <- val + 1
	}
}

func main() {
	c := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(2)

	go Even(&wg, c)
	go Odd(&wg, c)

	wg.Wait()
}
