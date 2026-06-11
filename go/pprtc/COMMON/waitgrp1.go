package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)
	fmt.Println("1")
	go printCat(&wg)
	wg.Wait()
	fmt.Println("2")
	go printDog(&wg)
	wg.Wait()

	fmt.Println("hello from main")
}

func printCat(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println("Cat")
	}
	wg.Done()

}

func printDog(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println("Dog")
	}
	wg.Done()
}
