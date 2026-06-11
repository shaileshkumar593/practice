package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aabbccdeaa"

	p := strings.Split(s, "")
	count := make(map[string]int)
	for _, x := range p {

		if v, found := count[x]; found {
			count[x] = v + 1
		} else {
			count[x] = 1
		}

	}
	for k, e := range count {
		fmt.Println(k, "----->", e)
	}
}
