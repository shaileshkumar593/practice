package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go count()
	time.Sleep(1)
	fmt.Println("end")
	for {
	}
}

func count() {
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		//time.Sleep(10)
	}
}
