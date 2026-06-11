package main

import "fmt"

func main() {
	mybyte := byte('b')

	myrune := '~'

	fmt.Printf("%c = %d  %c = %U\n", mybyte, mybyte, myrune, myrune)

}
