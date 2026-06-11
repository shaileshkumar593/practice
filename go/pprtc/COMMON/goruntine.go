package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("in main ")
	go func() {
		fmt.Println("inside annonimous goroutine ")
	}()
	fmt.Println("outside goroutine ")
	runtime.Gosched() // explicitly calling runtime go routine

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("hello second routine ")
	}(&wg)
	wg.Wait()

	fmt.Println("end of main")
}
