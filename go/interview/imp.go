package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func incrementCounter(wg *sync.WaitGroup, c chan bool) {
	defer wg.Done()

	for i := 0; i < 20; i++ {
		flag := <-c

		if flag == false {
			counter++
			c <- true
		}
	}
}

func printCounter(wg *sync.WaitGroup, c chan bool) {
	defer wg.Done()

	for i := 0; i < 20; i++ {
		flag := <-c

		if flag == true {
			fmt.Println(counter)
			c <- false
		}
	}
}

func main() {
	var wg sync.WaitGroup

	c := make(chan bool)

	wg.Add(2)

	go incrementCounter(&wg, c)
	go printCounter(&wg, c)

	// Give incrementCounter the first turn
	c <- false

	wg.Wait()
}
