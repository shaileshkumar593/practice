package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		//close(ch) // 🔥 mandatory
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
