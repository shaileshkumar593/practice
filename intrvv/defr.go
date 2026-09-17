package main

import (
	"fmt"
)

func Label() string {
	i := 0

	defer fmt.Println("deffered ", i)
	i = 5
	return fmt.Sprint("returned ", i)
}
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for {
		v, ok := <-ch
		fmt.Println(v, ok)
		if ok == false {
			break
		}
	}

	fmt.Println(Label())

}
