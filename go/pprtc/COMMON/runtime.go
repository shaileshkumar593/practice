package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup = sync.WaitGroup{}
var m sync.Mutex = sync.Mutex{}

func main() {
	//runtime.GOMAXPROCS(3)
	go func() {
		fmt.Println("Starting Annonimous first")

	}()
	fmt.Println("Below Annonimous first")
	runtime.Gosched()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Annonimous Second")
	}()
	go func() {
		fmt.Println("Starting Annonimous Third")

	}()
	time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println("Main is End here ")
	fmt.Println(" Max number of threads ", runtime.GOMAXPROCS(-1))
}
