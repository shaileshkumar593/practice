package main

import (
	"fmt"
)

func greet(name string, callback func(string)) {
	fmt.Println("greeting", name)

	callback(name)
}

func sayhello(s string) {
	fmt.Println("hello ", s)
}

func main() {
	greet("shailesh", sayhello)
}
