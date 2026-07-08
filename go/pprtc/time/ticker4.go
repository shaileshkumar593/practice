package main

import (
	"time"
	"fmt"
)

func main(){
	ticker := time.NewTicker(1*time.Second)
	defer ticker.Stop()
	var count int = 0
	for t:= range ticker.C{
		if count == 5{
			break
		}
		count = count + 1
		fmt.Println(t, " : ", count)
	}
}