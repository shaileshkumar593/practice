package main

import "fmt"

func main() {
	s := "Go😊"

	fmt.Println(len(s))    // 7 bytes (G=1, o=1, 😊=4)
	fmt.Println([]byte(s)) // [71 111 240 159 152 138]
	fmt.Println([]rune(s)) // [71 111 128522]
}
