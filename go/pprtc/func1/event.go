package main

import "fmt"

func onMessage(callback func(string)) {

	message := "New Event"

	callback(message)
}

func main() {

	onMessage(func(msg string) {
		fmt.Println("Received:", msg)
	})
}
