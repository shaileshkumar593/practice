package main

import (
	"fmt"
	"reflect"
)

type Author struct {
	name      string
	branch    string
	language  string
	particles int
}

func main() {

	a1 := Author{
		name:      "HC Verma",
		branch:    "Phy",
		language:  "Eng",
		particles: 45,
	}

	a2 := Author{
		name:      "HC Verma",
		branch:    "Phy",
		language:  "Eng",
		particles: 45,
	}

	a3 := Author{
		name:      "Dasgupta",
		branch:    "Math",
		language:  "Eng/Hindi",
		particles: 23,
	}

	fmt.Println(reflect.DeepEqual(a1, a2))
	fmt.Println(reflect.DeepEqual(a1, a3))
}
