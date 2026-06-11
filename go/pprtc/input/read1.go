package main

import (
	"fmt"
	"reflect"
)

func main() {
	const n int = 6
	//fmt.Scanf("%d\n", &n)
	var arr [n]int
	fmt.Println("Enter Value")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	for i, v := range arr {
		fmt.Println(i, v)
		fmt.Println(reflect.ValueOf(v).Kind())
		fmt.Println(reflect.TypeOf(v))
	}
}
