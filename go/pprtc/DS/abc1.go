package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 2, 5, 5, 4, 4, 7, 6, 8, 8} // {2,3,4,5,6,7,8}
	m := make(map[int]int, 0)

	for _, val := range s {
		if val, ok := m[val]; ok {
			m[val] = m[val] + 1
		} else {
			m[val] = 1
		}
	}

	fmt.Println(m)
	for key, _ := range m {
		fmt.Println(key)
	}
}
