package main

import (
	"fmt"
)

type image struct {
	data map[int]int
}

/*Struct variables are not comparable if they contain fields that are not comparable */
func main() {
	image1 := image{
		data: map[int]int{
			0: 155,
		}}
	image2 := image{
		data: map[int]int{
			0: 155,
		}}
	if image1 == image2 {
		fmt.Println("image1 and image2 are equal")
	}
}
