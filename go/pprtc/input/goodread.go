package main

import (
	"fmt"
)

func main() {
	var pp int               //
	var qq = make([]int, 10) //

	fmt.Println("Read  single value ")
	p, err := fmt.Scan(&pp)
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println("======", p)

	fmt.Println(" Reading Array of Int")

	for i := 0; i < len(qq); i++ {
		p, err = fmt.Scan(&qq[i])
		if err != nil {
			fmt.Println("Error ", err)
		}
		fmt.Println("------", p)
	}
}
