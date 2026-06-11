package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	defer func() {
		close(ch)
	}()
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println("Value of Channel i =", i)
		ch <- 13
		wg.Done()
	}()
	go func() {
		var i int
		i = 17
		ch <- i
		fmt.Println("Value of ch =", <-ch)
		wg.Done()
	}()
	wg.Wait()
}
