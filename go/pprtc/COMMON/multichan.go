package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func(ch chan string, msg string) {
		time.Sleep(2 * time.Second)
		ch <- fmt.Sprintf("Task 1 completed: %s", msg)
	}(chan1, "Hello how are you ")

	go func(ch chan string, msg string) {
		time.Sleep(4 * time.Second)
		ch <- fmt.Sprintf("Task2 completed %s", msg)
	}(chan2, "Welocome back")

	select {
	case msg1 := <-chan1:
		fmt.Println("chan1 message : ", msg1)

	case masg2 := <-chan2:
		fmt.Println("chan2 message: ", masg2)

	default:
		// Executes if no other case is ready
		fmt.Println("no channel is ready ") //  if default is added then output is no channel is ready
	}

}
