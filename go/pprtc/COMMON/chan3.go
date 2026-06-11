package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go countCat(c, &wg)

	for i := 0; i < 4; i++ {
		message := <-c
		fmt.Println(message)
	}
	wg.Wait()
}

func countCat(c chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		c <- "Cat"
	}
	wg.Done()
}
