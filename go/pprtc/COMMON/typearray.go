package main

import (
	"fmt"
	"strings"
)

func main() {
	var x interface{} = []int{1, 2, 3}
	fmt.Printf("Type : %T \nValue : %v\n\n", x, x)
	str := fmt.Sprintf("%v", x)
	fmt.Println(str)
	fmt.Printf("Type : %T \nValue : %v\n\n", str, str)

	var y interface{} = []string{"a", "2", "3"}
	fmt.Printf("Type : %T \nValue : %v\n\n", y, y)
	str1 := fmt.Sprintf("%v", y)
	fmt.Println(str1)
	fmt.Printf("Type : %T \nValue : %v\n\n", str1, str1)
	for _, val := range str1 {
		fmt.Println(val)
	}
	fmt.Println("len :", len(str1))
	//fmt.Println(strings.Join(str1, ":"))  str1 type is string not []string
	str3 := strings.Fields(str1)
	fmt.Printf("Type : %T \nValue : %v\n\n", str3, str3s)
	s1 := "abc"
	fmt.Printf("Type : %T \nValue : %v\n\n", s1, s1)
	str2 := fmt.Sprintf("%v", s1)
	fmt.Println(str2)
	fmt.Printf("Type : %T \nValue : %v\n\n", str2, str2)
}
