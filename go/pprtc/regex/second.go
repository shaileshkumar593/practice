package main

import (
	"fmt"
	"net/mail"
)

func validMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

var addresses = []string{
	"foo@gmail.compppppppppppppppp",
	"Gopher <from@example.com>",
	"example",
}

func main() {
	for _, a := range addresses {
		if addr, ok := validMailAddress(a); ok {
			fmt.Printf("value: %-30s valid email: %-10t address: %s\n", a, ok, addr)
		} else {
			fmt.Printf("value: %-30s valid email: %-10t\n", a, ok)
		}
	}
}
