package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)
	go A(&wg)
	go B(&wg)

	wg.Wait()

}
func A(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 5) //io operation delay
	fmt.Println("Function A")
}
func B(wg *sync.WaitGroup) {
	defer wg.Done()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("unrechable")
		}
	}()
	panic("unreachable")

}
