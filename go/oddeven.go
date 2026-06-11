package main

import (
	"fmt"
	"sync"
)

func OddCount(wg *sync.WaitGroup, oddch, evench chan struct{}) {
	defer wg.Done()
	for i := 1; i <= 9; i += 2 {
		<-oddch
		fmt.Println(i)
		evench <- struct{}{}

	}
}

func EvenCount(wg *sync.WaitGroup, oddch, evench chan struct{}) {
	defer wg.Done()

	for i := 0; i <= 10; i += 2 {
		<-evench
		fmt.Println(i)

		if i != 10 {
			oddch <- struct{}{}
		}

	}
}

func main() {
	oc := make(chan struct{})
	ec := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(2)

	go OddCount(&wg, oc, ec)
	go EvenCount(&wg, oc, ec)

	oc <- struct{}{}

	wg.Wait()

}
