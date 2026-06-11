package main

import (
	"fmt"
	"sync"
)

func Even(wg *sync.WaitGroup, evenCh, oddCh chan int) {
	defer wg.Done()

	val := 0
	fmt.Println(val)

	for {
		val++

		oddCh <- val

		val = <-evenCh

		if val == 10 {
			fmt.Println(val)
			close(oddCh)
			return
		}
		fmt.Println(val)
	}
}

func Odd(wg *sync.WaitGroup, evenCh, oddCh chan int) {
	defer wg.Done()

	for {
		val, ok := <-oddCh

		if !ok {
			close(evenCh)
			return
		}

		fmt.Println(val)

		evenCh <- val + 1
	}
}

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(2)

	go Even(&wg, evenCh, oddCh)
	go Odd(&wg, evenCh, oddCh)

	wg.Wait()
}
