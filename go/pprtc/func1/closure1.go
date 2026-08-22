package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {

		func() {

			fmt.Println(i)

		}()

	}
}
