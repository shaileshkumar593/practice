package main

import (
	"fmt"
	"time"
)

func Counter(c chan int) {
	a := <-c
	b := <-c
	close(c)
	fmt.Println(a * b)
}
func main() {
	var ch = make(chan int)

	go Counter(ch)
	ch <- 26
	ch <- 28
	_, status := <-ch
	fmt.Println(status) // false channel close
	time.Sleep(1 * time.Second)
	fmt.Println("End of main")

}
