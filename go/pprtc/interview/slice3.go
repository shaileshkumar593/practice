package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// 🔹 41. Slice of structs
func sliceOfStructs() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}
	fmt.Println("People:", people)
}

// 🔹 42. Passing slice to function
func sumSlice(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// 🔹 43. Returning slice from function
func evenNumbers(limit int) []int {
	var result []int
	for i := 0; i <= limit; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	return result
}

// 🔹 44. Slice of slices (ragged)
func raggedSlice() {
	matrix := [][]int{
		{1, 2},
		{3, 4, 5},
		{6},
	}
	fmt.Println("Ragged:", matrix)
}

// 🔹 45. Deep copy vs shallow copy
func deepShallow() {
	orig := []int{1, 2, 3}
	copySlice := orig                    // shallow copy (shares memory)
	indep := append([]int(nil), orig...) // deep copy
	orig[0] = 99
	fmt.Println("Shallow:", copySlice)
	fmt.Println("Deep:", indep)
}

// 🔹 46. Slicing strings
func stringSlice() {
	s := "golang"
	fmt.Println("Substring:", s[0:3]) // "gol"
}

// 🔹 47. Modify string using []rune slice
func modifyString() {
	s := "hello"
	r := []rune(s)
	r[0] = 'H'
	fmt.Println("Modified:", string(r))
}

// 🔹 48. Slice capacity growth
func sliceGrowth() {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("Len=%d Cap=%d\n", len(s), cap(s))
	}
}

// 🔹 49. Slice in JSON
func sliceJSON() {
	data := []string{"go", "java", "rust"}
	j, _ := json.Marshal(data)
	fmt.Println("JSON:", string(j))
}

// 🔹 50. JSON → slice
func jsonToSlice() {
	jsonStr := `["cat","dog","bird"]`
	var animals []string
	_ = json.Unmarshal([]byte(jsonStr), &animals)
	fmt.Println("Unmarshaled:", animals)
}

// 🔹 51. Slice as function argument modifies original
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 100
	}
}

// 🔹 52. Slice copy doesn’t modify original
func safeModify(s []int) {
	c := append([]int(nil), s...)
	c[0] = 200
	fmt.Println("Inside safeModify:", c)
}

// 🔹 53. Slice of maps
func sliceOfMaps() {
	users := []map[string]int{
		{"Alice": 25},
		{"Bob": 30},
	}
	fmt.Println("Slice of maps:", users)
}

// 🔹 54. Map of slices
func mapOfSlices() {
	courses := map[string][]string{
		"CS": {"Go", "Java"},
		"EE": {"Circuits", "Signals"},
	}
	fmt.Println("Map of slices:", courses)
}

// 🔹 55. Append inside loop (careful with capacity reuse)
func appendInsideLoop() {
	s := []int{}
	for i := 0; i < 5; i++ {
		s = append(s, i)
	}
	fmt.Println("After loop:", s)
}

// 🔹 56. Remove duplicates
func unique(nums []int) []int {
	seen := make(map[int]bool)
	var res []int
	for _, v := range nums {
		if !seen[v] {
			res = append(res, v)
			seen[v] = true
		}
	}
	return res
}

// 🔹 57. Slice of interfaces
func sliceOfInterfaces() {
	var any []interface{} = []interface{}{1, "hello", 3.14}
	fmt.Println("Interfaces slice:", any)
}

// 🔹 58. Slice re-slicing
func reSlice() {
	arr := []int{10, 20, 30, 40, 50}
	s1 := arr[1:4]
	s2 := s1[:2]
	fmt.Println("s1:", s1, "s2:", s2)
}

// 🔹 59. Slice alias problem
func aliasProblem() {
	arr := []int{1, 2, 3}
	a := arr[0:2]
	b := arr[1:3]
	a[1] = 99
	fmt.Println("Arr:", arr, "B:", b)
}

// 🔹 60. Slice capacity control
func capControl() {
	arr := []int{1, 2, 3, 4, 5}
	s := arr[:2:2] // len=2 cap=2
	fmt.Println("Len:", len(s), "Cap:", cap(s))
}

// 🔹 61. Growing slice with doubling
func growSlice() {
	s := make([]int, 0, 1)
	for i := 0; i < 8; i++ {
		s = append(s, i)
		fmt.Printf("%d => Len=%d Cap=%d\n", i, len(s), cap(s))
	}
}

// 🔹 62. Slice comparison (manual, since Go doesn’t support ==)
func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 🔹 63. Passing nil slice to function
func nilSliceFunc(s []int) {
	if s == nil {
		fmt.Println("Got nil slice")
	}
}

// 🔹 64. Variadic function → slice
func variadic(nums ...int) {
	fmt.Println("Received:", nums)
}

// 🔹 65. Flatten slice of slices
func flatten(matrix [][]int) []int {
	var flat []int
	for _, row := range matrix {
		flat = append(flat, row...)
	}
	return flat
}

// 🔹 66. Slice as stack
func stack() {
	var st []int
	st = append(st, 10)
	st = append(st, 20)
	fmt.Println("Stack:", st)
	st = st[:len(st)-1]
	fmt.Println("Pop:", st)
}

// 🔹 67. Slice as queue
func queue() {
	var q []int
	q = append(q, 1, 2, 3)
	fmt.Println("Queue:", q)
	q = q[1:]
	fmt.Println("After dequeue:", q)
}

// 🔹 68. Rotate slice left
func rotateLeft(s []int, k int) []int {
	k %= len(s)
	return append(s[k:], s[:k]...)
}

// 🔹 69. Rotate slice right
func rotateRight(s []int, k int) []int {
	k %= len(s)
	return append(s[len(s)-k:], s[:len(s)-k]...)
}

// 🔹 70. Prefix sum with slice
func prefixSum(nums []int) []int {
	prefix := make([]int, len(nums))
	sum := 0
	for i, v := range nums {
		sum += v
		prefix[i] = sum
	}
	return prefix
}

func main() {
	sliceOfStructs()
	fmt.Println("Sum:", sumSlice([]int{1, 2, 3}))
	fmt.Println("Evens:", evenNumbers(10))
	raggedSlice()
	deepShallow()
	stringSlice()
	modifyString()
	sliceGrowth()
	sliceJSON()
	jsonToSlice()

	nums := []int{1, 2, 3}
	modifySlice(nums)
	fmt.Println("After modifySlice:", nums)
	safeModify(nums)

	sliceOfMaps()
	mapOfSlices()
	appendInsideLoop()
	fmt.Println("Unique:", unique([]int{1, 2, 2, 3, 3, 4}))
	sliceOfInterfaces()
	reSlice()
	aliasProblem()
	capControl()
	growSlice()
	fmt.Println("Compare:", compareSlices([]int{1, 2}, []int{1, 2}))
	nilSliceFunc(nil)
	variadic(10, 20, 30)

	matrix := [][]int{{1, 2}, {3, 4}}
	fmt.Println("Flatten:", flatten(matrix))
	stack()
	queue()
	fmt.Println("Rotate Left:", rotateLeft([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println("Rotate Right:", rotateRight([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println("Prefix sum:", prefixSum([]int{1, 2, 3, 4}))
}
