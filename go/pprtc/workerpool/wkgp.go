package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)
	var wg *sync.WaitGroup
	wg.Add(1)

	go countCat(c, wg)
	c <- "car"

	wg.Wait()
}

func countCat(c chan string, wg *sync.WaitGroup) {
	msg := <-c
	fmt.Println(msg)
	wg.Done()
}
