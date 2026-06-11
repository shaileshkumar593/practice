package main

import (
	"fmt"
	"strconv"
)

func main() {

	no := "101"

	s, err := strconv.ParseInt(no, 10, 32)

	if err != nil {

		fmt.Print(err)

	}

	fmt.Printf("%T, %v\n", s, s)

	s, err = strconv.ParseInt(no, 16, 32)

	if err != nil {

		fmt.Print(err)

	}

	fmt.Printf("%T, %v\n", s, s)

	s, err = strconv.ParseInt(no, 10, 64)

	if err != nil {

		fmt.Print(err)

	}

	fmt.Printf("%T, %v\n", s, s)

	s, err = strconv.ParseInt(no, 16, 64)

	if err != nil {

		fmt.Print(err)
	}

	fmt.Printf("%T, %v\n", s, s)

}
