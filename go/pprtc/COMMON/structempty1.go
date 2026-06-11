package main

import (
	"fmt"
	"reflect"
)

type Vehicle struct {
	size int
}

func (x Vehicle) IsStructureEmpty() bool {

	return reflect.DeepEqual(x, Vehicle{})

}

func main() {

	vehicle := Vehicle{}

	if vehicle.IsStructureEmpty() {

		fmt.Println("It is empty structure")

	} else {

		fmt.Println("It is not empty structure")

	}
}
