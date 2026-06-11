package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Counter(c chan int) {
	a := <-c
	b := <-c
	close(c)
	fmt.Println(a * b)
	wg.Done()
}
func main() {
	var ch = make(chan int)
	wg.Add(1)
	go Counter(ch)
	wg.Wait()
	ch <- 26
	ch <- 28
	_, status := <-ch
	fmt.Println(status) // false channel close
	//time.Sleep(1 * time.Second)
	fmt.Println("End of main")

}
