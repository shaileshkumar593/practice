package main

import "fmt"

func swap(i, j *int) {
	temp := *i
	*i = *j
	*j = temp
}
func main() {
	a, b := 22, 77
	fmt.Println(a, "----------->", b)

	swap(&a, &b)
	fmt.Println(a, "----------->", b)
}
