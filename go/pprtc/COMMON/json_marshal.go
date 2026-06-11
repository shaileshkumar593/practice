package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	var x interface{} = []int{1, 2, 3}
	fmt.Printf("Type : %T \nValue : %v\n\n", x, x)

	jsonStr, err := json.Marshal(x)
	fmt.Printf("Type : %T \nValue : %v\n\n", jsonStr, jsonStr)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())

	} else {
		fmt.Println(string(jsonStr))
		fmt.Printf("Type : %T \nValue : %v\n\n", string(jsonStr), string(jsonStr))
	}

	fmt.Printf("Type : %T \nValue : %v\n\n",_,strings.Split(string(jsonStr,""))

}
