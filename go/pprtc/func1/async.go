package main

import (
	"fmt"
	"time"
)

func fetchData(callback func(string)) {

	go func() {

		time.Sleep(2 * time.Second)

		callback("Data Loaded")
	}()
}

func main() {

	fetchData(func(result string) {
		fmt.Println(result)
	})

	fmt.Println("Main continues...")

	time.Sleep(3 * time.Second)
}
