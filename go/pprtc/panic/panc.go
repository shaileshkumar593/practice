package main

import (
	"fmt"
)

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

/*
	deferred call in fullName
deferred call in main
panic: runtime error: last name cannot be nil

goroutine 1 [running]:
main.fullName(0x60?, 0xc000054000?)
        D:/shail/poc/goconcurrency/panic/panc.go:13 +0x177
main.main()
        D:/shail/poc/goconcurrency/panic/panc.go:22 +0x85
exit status 2
*/
