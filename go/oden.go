package main

import (
	"fmt"
	"sync"
)

func printOdd(wg *sync.WaitGroup, oddCh, evenCh chan struct{}) {
	defer wg.Done()

	for i := 1; i <= 9; i += 2 {
		<-oddCh
		fmt.Println(i)
		evenCh <- struct{}{}
	}
}

func printEven(wg *sync.WaitGroup, oddCh, evenCh chan struct{}) {
	defer wg.Done()

	for i := 0; i <= 10; i += 2 {
		<-evenCh
		fmt.Println(i)

		// avoid extra send after last
		if i != 10 {
			oddCh <- struct{}{}
		}
	}
}

func main() {
	var wg sync.WaitGroup

	oddCh := make(chan struct{})
	evenCh := make(chan struct{})

	wg.Add(2)

	go printOdd(&wg, oddCh, evenCh)
	go printEven(&wg, oddCh, evenCh)

	// 🔥 START WITH EVEN (key change)
	evenCh <- struct{}{}

	wg.Wait()
}
