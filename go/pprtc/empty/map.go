package main

import "fmt"

func main() {
	map_create := make(map[string]int)
	if len(map_create) == 0 {
		fmt.Println("Map is empty")
	} else {
		fmt.Println("Map is not empty")
	}
}
