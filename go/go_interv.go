package main

import (
	"fmt"
	"reflect"
)

type S struct {
	A string
	B string
	C string
}

func main() {

	s := "123"

	ps := &s

	b := []byte(*ps)

	pb := &b

	s += "4"

	*ps += "5"

	b[1] = '0'

	fmt.Println(*ps)

	fmt.Println(string(*pb))

	x := interface{}(&S{"a", "b", "c"})

	y := interface{}(&S{"a", "b", "c"})

	fmt.Println(x == y)

	fmt.Println(reflect.DeepEqual(x, y))

}
