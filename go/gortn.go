package main

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	//val := <-ch

	//fmt.Println(val)
}

/*

go routine blocked   but main has nothing to execute so execution runs */
