package main

import "fmt"

func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}
func main() {
	defer second()
	first()
	fmt.Println("hello main")
}

/*Deferred Functions are run even if a runtime panic occurs.
Deferred functions are executed in LIFO order, so the above code prints: 4 3 2 1 0.
You can put multiple functions on the "deferred list", like this example.
*/
