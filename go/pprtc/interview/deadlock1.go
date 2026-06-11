package main

func main() {
	c := make(chan int)

	<-c // dead lock due to channel
}
