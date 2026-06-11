package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("%d", i)
		}()
	}
	time.Sleep(time.Second)
}
