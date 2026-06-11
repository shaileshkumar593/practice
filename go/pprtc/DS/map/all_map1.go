package main

import "fmt"

// Custom Types
type User23 struct {
	Name string
	Age  int
}

type Point struct {
	X int
	Y int
}

func main() {

	fmt.Println("======== BASIC KEY TYPES ========")

	// int as key
	var m1 map[int]string = map[int]string{1: "one"}
	fmt.Println("int -> string:", m1)

	m2 := map[int][]int{10: {1, 2, 3}}
	fmt.Println("int -> slice:", m2)

	m3 := map[int]map[string]int{1: {"a": 10}}
	fmt.Println("int -> map:", m3)

	m4 := map[int]User23{1: {"Alice", 25}}
	fmt.Println("int -> struct:", m4)

	m5 := map[int]func(int) int{1: func(x int) int { return x * 2 }}
	fmt.Println("int -> function:", m5)

	m6 := map[int]interface{}{2: "hello"}
	fmt.Println("int -> interface{}:", m6)

	m7 := map[int]*User23{1: &User{"Bob", 30}}
	fmt.Println("int -> pointer:", m7)

	m8 := map[int]chan int{1: make(chan int)}
	fmt.Println("int -> channel:", m8)

	// ------------------------------------

	// string as key
	m9 := map[string]int{"x": 100}
	fmt.Println("string -> int:", m9)

	m10 := map[string][]string{"names": {"sam", "tom"}}
	fmt.Println("string -> slice:", m10)

	m11 := map[string]map[int]bool{"outer": {1: true}}
	fmt.Println("string -> map:", m11)

	m12 := map[string]User{"u1": {"John", 20}}
	fmt.Println("string -> struct:", m12)

	m13 := map[string]func(string) string{"hello": func(s string) string { return "Hi, " + s }}
	fmt.Println("string -> func:", m13["hello"]("Go"))

	m14 := map[string]interface{}{"key": 55.5}
	fmt.Println("string -> interface:", m14)

	m15 := map[string]*User23{"bob": {Name: "Bob", Age: 35}}
	fmt.Println("string -> pointer:", m15)

	m16 := map[string]chan bool{"done": make(chan bool)}
	fmt.Println("string -> channel:", m16)

	// ------------------------------------

	fmt.Println("\n======== STRUCT AS KEY ========")

	p1 := Point{X: 1, Y: 2}

	m17 := map[Point]string{{1, 2}: "point-1"}
	fmt.Println("struct -> string:", m17[p1])

	m18 := map[Point][]int{{1, 2}: {10, 20}}
	fmt.Println("struct -> slice:", m18)

	m19 := map[Point]map[string]int{{1, 2}: {"age": 30}}
	fmt.Println("struct -> map:", m19)

	m20 := map[Point]User23{{1, 2}: {"Max", 40}}
	fmt.Println("struct -> struct:", m20)

	m21 := map[Point]interface{}{{1, 2}: "data"}
	fmt.Println("struct -> interface:", m21)

	// ------------------------------------

	fmt.Println("\n======== ARRAY AS KEY ========")

	m22 := map[[3]int]string{{1, 2, 3}: "array"}
	fmt.Println("array -> string:", m22)

	m23 := map[[2]string]int{{"a", "b"}: 100}
	fmt.Println("array -> int:", m23)

	m24 := map[[4]byte]User23{{1, 2, 3, 4}: {"Zed", 50}}
	fmt.Println("array -> struct:", m24)

	// ------------------------------------

	fmt.Println("\n======== POINTER AS KEY ========")

	a := 10
	m25 := map[*int]string{&a: "value"}
	fmt.Println("pointer -> string:", m25[&a])

	u := &User{"PointerUser", 44}
	m26 := map[*User]int{u: 500}
	fmt.Println("pointer -> int:", m26[u])

	m27 := map[*User]map[int]bool{u: {1: true}}
	fmt.Println("pointer -> map:", m27)

	// ------------------------------------

	fmt.Println("\n======== CHANNEL AS KEY ========")

	ch1 := make(chan int)
	m28 := map[chan int]string{ch1: "channelKey"}
	fmt.Println("chan -> string:", m28)

	ch2 := make(chan bool)
	m29 := map[chan bool]int{ch2: 999}
	fmt.Println("chan -> int:", m29)

	// ------------------------------------

	fmt.Println("\n======== INTERFACE AS KEY (value must be comparable) ========")

	m30 := map[interface{}]string{
		"a":  "alpha",   // string
		1:    "one",     // int
		true: "boolKey", // bool
	}
	fmt.Println("interface -> string:", m30)

	// interface value types
	m31 := map[int]interface{}{
		1: "hello",
		2: 10,
		3: []int{1, 2, 3}, // ok: slice as VALUE
	}
	fmt.Println("int -> interface:", m31)

	// ------------------------------------

	fmt.Println("\n======== FUNCTION AS VALUE ========")

	m32 := map[string]func(int) int{
		"square": func(n int) int { return n * n },
	}
	fmt.Println("function value:", m32)

	// ------------------------------------

	fmt.Println("\n======== CUSTOM TYPE KEY ========")

	type ID int
	m33 := map[ID]string{1: "custom"}
	fmt.Println("custom key -> string:", m33)

	type Email string
	m34 := map[Email]int{"test@example.com": 1}
	fmt.Println("custom key -> int:", m34)
}
