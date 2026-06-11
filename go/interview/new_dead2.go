package main

import (
	"fmt"
	"time"
)

func Abc(c chan int) {
	fmt.Println(<-c)
}
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 5; i++ {
		ch <- i
		go Abc(ch)

	}
	time.Sleep(5 * time.Second)
}

// output
/*
	PS D:\concurrency\interview> go run .\new_dead2.go
0
1
2
3
4
*/
