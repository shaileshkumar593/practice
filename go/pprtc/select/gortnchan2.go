package main

import (
	"fmt"
)

func PrintHello(p chan string) {

	fmt.Println("Hello How are you", <-p)
	fmt.Println("In PrintHello go routine ")
}

func main() {
	c := make(chan string)
	go PrintHello(c)
	c <- "shailesh" // here wait on channel till data receive

	fmt.Println("In Main go routine ")

}

/*PS D:\shail\poc\goconcurrency\select> go run .\gortnchan2.go
Hello How are you shailesh
In PrintHello go routine
In Main go routine*/
