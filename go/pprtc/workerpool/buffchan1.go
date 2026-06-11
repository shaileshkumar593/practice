package main

func main() {
	ch := make(chan string, 0) //make(chan string)
	ch <- "naveen"
	//ch <- "paul"
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	// deadlock error issue
}
