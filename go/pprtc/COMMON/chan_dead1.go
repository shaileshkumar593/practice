package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	go dead_lock(c)
	for i := 0; i < 10; i++ {
		c <- "Bye"
	}
}

/* sender is sending 10 msg butreceiver is receiving only 9 here deadlock happens"
 */
func dead_lock(c1 chan string) {
	for i := 0; i < 9; i++ {
		msg, _ := <-c1
		/* if ok {
			fmt.Println(msg)  handling channel issue
		} */
		fmt.Println(msg)
	}
}
