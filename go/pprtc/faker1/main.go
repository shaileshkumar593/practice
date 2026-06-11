package main

import (
	"fmt"

	"github.com/jaswdr/faker"
)

func main() {
	faker := faker.New()
	fmt.Println(faker.Person().Name())
}
