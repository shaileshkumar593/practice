package main

import "fmt"

func firstOccurence(input string) string {

	m := make(map[rune]struct{})

	rslt := make([]rune, 0)

	for _, val := range input {
		if _, exist := m[val]; !exist {
			m[val] = struct{}{}

			rslt = append(rslt, val)
		}
	}

	return string(rslt)
}

func main() {
	input := "programing"

	fmt.Println(firstOccurence(input))
}
