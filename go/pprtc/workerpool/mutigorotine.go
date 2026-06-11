package main

import (
	"fmt"
	"sync"
)

func printinfo(c chan int) {

	val, _ := <-c
	fmt.Println(val)
}
func main() {
	var wg sync.WaitGroup
	c := make(chan int, 2)
	wg.Add(1)

	go printinfo(c)
	go printinfo(c)
	for i := 0; i < 7; i++ {

		c <- i

	}

	//
	wg.Done()

	wg.Wait()

}
