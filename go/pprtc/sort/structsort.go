package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Marks int
}

type ByMarks []Student

func (s ByMarks) Len() int {
	return len(s)
}

func (s ByMarks) Less(i, j int) bool {
	return s[i].Marks < s[j].Marks
}

/*
	func (s ByMarks) Less(i, j int) bool {
	if s[i].Marks == s[j].Marks {
		return s[i].Name < s[j].Name
	}
	return s[i].Marks < s[j].Marks
}

*/


func (s ByMarks) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	students := []Student{
		{"A", 80},
		{"B", 95},
		{"C", 70},
	}

	sort.Sort(ByMarks(students))

	fmt.Println(students)
}