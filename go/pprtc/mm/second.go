package main

import (
	"fmt"

	"github.com/jaswdr/faker"
)

func main() {

	f := faker.New()
	type ExampleStruct struct {
		SimpleStringField  string
		SimpleNumber       int
		SimpleBool         bool
		SomeFormatedString string    `fake:"??? ###"`
		SomeStringArray    [5]string `fake:"????"`
	}

	example := ExampleStruct{}
	f.Struct().Fill(&example)
	fmt.Printf("%+v", example)
}
