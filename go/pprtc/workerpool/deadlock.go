package main

func main() {
	ch := make(chan int)
	ch <- 5
	// if write on channel then there must be
	//some read happen otherwise channel is blocked and cause deadlock
}
