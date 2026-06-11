package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	fmt.Println("\n================= SLICES =================")

	// 1. Nil slice
	var s1 []int
	fmt.Println("Nil slice:", s1, "len:", len(s1), "cap:", cap(s1))
	s1 = append(s1, 10) // OK
	fmt.Println("After append:", s1)

	// 2. Empty slice literal
	s2 := []int{}
	fmt.Println("Empty literal slice:", s2)

	// 3. Slice with initial values
	s3 := []int{1, 2, 3}
	fmt.Println("Slice with values:", s3)

	// 4. Slice using make
	s4 := make([]int, 3, 5)
	fmt.Println("Make slice:", s4, "len:", len(s4), "cap:", cap(s4))

	// 5. Slice from array
	arr := [5]int{10, 20, 30, 40, 50}
	s5 := arr[1:4] // 20,30,40
	fmt.Println("Slice from array:", s5)

	// 6. Slice from slice
	s6 := s3[1:3] // 2,3
	fmt.Println("Slice from slice:", s6)

	// 7. Full slice expression
	s7 := arr[1:3:4]
	fmt.Println("Full slice:", s7, "len:", len(s7), "cap:", cap(s7))

	// 8. Slice of structs
	users := []User{
		{"Alice", 25},
		{"Bob", 30},
	}
	fmt.Println("Slice of structs:", users)

	// 9. Slice of interfaces (mixed types)
	var s9 []interface{}
	s9 = append(s9, "hello", 10, 5.5)
	fmt.Println("Slice of interface:", s9)

	// 10. 2D Slice (matrix)
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("2D Slice:", matrix)

	fmt.Println("\n================= MAPS =================")

	// 1. Nil map
	var m1 map[string]int
	fmt.Println("Nil map:", m1)
	// m1["x"] = 10 // panic: assignment to nil map

	// 2. Empty map using make
	m2 := make(map[string]int)
	m2["a"] = 10
	fmt.Println("Map from make:", m2)

	// 3. Map literal
	m3 := map[string]int{"x": 1, "y": 2}
	fmt.Println("Map literal:", m3)

	// 4. Map with capacity
	m4 := make(map[int]string, 10)
	m4[1] = "One"
	fmt.Println("Map with capacity:", m4)

	// 5. Map of slices
	m5 := map[string][]int{
		"numbers": {1, 2, 3},
	}
	m5["numbers"] = append(m5["numbers"], 4)
	fmt.Println("Map of slices:", m5)

	// 6. Map of structs
	m6 := map[string]User{
		"user1": {"John", 22},
		"user2": {"Mira", 30},
	}
	fmt.Println("Map of structs:", m6)

	// 7. Map of maps
	m7 := map[string]map[string]int{
		"outer": {"inner": 100},
	}
	m7["outer"]["inner"] = 200
	fmt.Println("Map of maps:", m7)

	// 8. Map with interface values
	m8 := map[string]interface{}{
		"name": "Shailesh",
		"age":  27,
		"gpa":  8.4,
	}
	fmt.Println("Map with interface values:", m8)

	// 9. Map of functions
	m9 := map[string]func(int) int{
		"square": func(x int) int { return x * x },
		"double": func(x int) int { return x * 2 },
	}
	fmt.Println("Function map output: square(5) =", m9)
	fmt.Println("Function map output: double(5) =", m9)

	// 10. Map with custom type key
	type ID int
	m10 := map[ID]string{
		1: "Hello",
		2: "World",
	}
	fmt.Println("Custom type map:", m10)
}
