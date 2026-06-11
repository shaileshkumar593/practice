package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var m sync.Mutex

func main() {
	c := make(chan string)
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go send_msg(c)

		msg, ok := <-c
		if ok {
			fmt.Println(msg)
		}

	}
	wg.Wait()
}

// with and without lock it works
func send_msg(c chan string) {
	m.Lock()
	c <- "Hello"
	m.Unlock()
	//close(c)// panic error
	wg.Done()

}
