package main

import (
	"fmt"
	"time"
)

func PrintHello(p chan string) {

	fmt.Println("Hello How are you", <-p)
	time.Sleep(2 * time.Second)
	fmt.Println("In PrintHello go routine ")
}

func main() {
	c := make(chan string)
	go PrintHello(c)
	c <- "shailesh" // here wait on channel till data receive

	fmt.Println("In Main go routine ")

}

/*
PS D:\shail\poc\goconcurrency\select> go run .\gortnchan3.go
Hello How are you shailesh
In Main go routine
*/
