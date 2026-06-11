package main

import (
	"fmt"
	"sync"
)

func odd(oddCh, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 9; i += 2 {
		<-oddCh
		fmt.Println(i)
		evenCh <- 1
	}
}

func even(oddCh, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 2; i <= 10; i += 2 {
		<-evenCh
		fmt.Println(i)

		if i != 10 {
			oddCh <- 1
		}
	}
}

func main() {
	var wg sync.WaitGroup

	oddCh := make(chan int)
	evenCh := make(chan int)

	wg.Add(2)

	go odd(oddCh, evenCh, &wg)
	go even(oddCh, evenCh, &wg)

	oddCh <- 1

	wg.Wait()
}
