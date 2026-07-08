package main

import (
	"fmt"
	"sort"
	"strconv"
)

type ValueConv struct {
	s   string
	val int
}

func main() {

	str := []string{"i20", "c5", "i10", "T2", "T1", "T15"}

	arr := make([]ValueConv, len(str))

	for i, s := range str {

		num, err := strconv.Atoi(s[1:])
		if err != nil {
			panic(err)
		}

		arr[i] = ValueConv{
			s:   s,
			val: num,
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].val < arr[j].val
	})

	for i := range arr {
		str[i] = arr[i].s
	}

	fmt.Println(str)
}