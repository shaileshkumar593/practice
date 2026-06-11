package main

import "fmt"

type person struct {
	fname string
	lname string
}

type employee struct {
	person
	empid int
}

func (p person) details() {
	fmt.Println(p, " "+"I am a person")
}

func (e employee) details() {
	fmt.Println(e, " "+"I am a Employee")
}

func main() {
	p1 := person{"Raj", "Kumar"}
	p1.details()
	e1 := employee{person: person{"john", "Kumar"}, empid: 23}
	e1.details()
}
