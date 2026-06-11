package main

import "fmt"

type Store struct {
	Index int
	Count int
}

var storetype Store

func Non_repeat(s string) {

	ps := []rune(s)
	m := make(map[rune]Store)

	for index, runeval := range ps {
		if val, exist := m[runeval]; exist {
			m[runeval] = Store{
				Count: val.Count + 1,
				Index: index,
			}
		} else {
			m[runeval] = Store{
				Count: 1,
				Index: index,
			}
		}
	}

	for key, val := range m {
		if val.Count == 1 {
			fmt.Println(string(key), " : ", val.Index)
		}
	}

}
func main() {
	input := "loveleetcode"

	Non_repeat(input)
}
