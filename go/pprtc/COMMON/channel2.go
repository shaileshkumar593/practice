/*Actually, we can pull out two variables from the channel, one is the value
and the other is a bool value (comma OK) which tells that the channel is open
or not. If the channel is open the bool value returns true and false if the
channel is closed.*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		var i int
		i = 17
		ch <- i
		ch <- i + 18
		ch <- 13
		ch <- 19
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

/*Actually
go func(ch chan<- int) {
                defer func() {
                    close(ch)
                }()
		var i int
		i = 17
		ch <- i
		ch <- i + 18
		ch <- 13
		ch <- 19
		wg.Done()
	}(ch)
	wg.Wait()*/
