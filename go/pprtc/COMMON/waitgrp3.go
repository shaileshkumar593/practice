package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	//wg1 := sync.WaitGroup{}
	wg.Add(2)

	fmt.Println("1")
	go printCat(&wg)
	fmt.Println("2")
	//wg.Wait()

	//wg1.Add(1)
	fmt.Println("3")
	go printDog(&wg)
	fmt.Println("4")
	wg.Wait()
	fmt.Println("5")

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

// progm runs
