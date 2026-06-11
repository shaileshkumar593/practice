package main

import "fmt"

func PrintHello(p chan string) {
	fmt.Println("In PrintHello go routine ")
	fmt.Println("Hello How are you", <-p)
}

func main() {
	c := make(chan string)
	go PrintHello(c)
	c <- "shailesh" // here wait on channel till data receive

	fmt.Println("In Main go routine ")

}
