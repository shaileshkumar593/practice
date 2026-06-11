package main

import (
	"fmt"
	"sort"
)

func findTheDifferenceSort(s, t string) byte {
	sArr := []byte(s)
	tArr := []byte(t)

	sort.Slice(sArr, func(i, j int) bool { return sArr[i] < sArr[j] })
	sort.Slice(tArr, func(i, j int) bool { return tArr[i] < tArr[j] })

	for i := 0; i < len(sArr); i++ {
		if sArr[i] != tArr[i] {
			return tArr[i]
		}
	}

	return tArr[len(tArr)-1]
}

func main() {
	fmt.Println(string(findTheDifferenceSort("abcd", "ceadb")))
}
