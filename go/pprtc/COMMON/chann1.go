package main

import (
	"fmt"
)

// study GoRoutine
// Main GoRoutine and all other GoRoutine run concurrently
// So once the main goroutine termintes then the execution of programme stops.

func main() {

	fmt.Println("We are in main start")
	go myfunc()

	fmt.Println("We are the End of main")

}

func myfunc() {
	fmt.Println("We are  in ***GoRoutine********")
}
