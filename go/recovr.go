package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				fmt.Println("Error:", err)
			}
		}
	}()

	panic(errors.New("worker failed"))
}
