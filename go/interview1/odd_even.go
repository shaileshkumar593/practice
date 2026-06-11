package main

import (
	"fmt"
	"sync"
)

func printOdd(wg *sync.WaitGroup, oddCh, evenCh chan struct{}) {
	defer wg.Done()

	for i := 1; i <= 9; i += 2 {
		<-oddCh              // wait for turn
		fmt.Println(i)       // print odd
		evenCh <- struct{}{} // signal even
	}
}

func printEven(wg *sync.WaitGroup, oddCh, evenCh chan struct{}) {
	defer wg.Done()

	for i := 0; i <= 10; i += 2 {
		<-evenCh       // wait for turn
		fmt.Println(i) // print even
		if i != 10 {
			oddCh <- struct{}{} // signal odd
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

	// 🔥 start with odd
	oddCh <- struct{}{}

	wg.Wait()
}
