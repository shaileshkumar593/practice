package main

import (
	"fmt"
	"sync"
)

int i =0 
var odd,even bool
var c chan bool = false
func oddPrint(){
	if odd{
		fmt.Println(i++)
		even=true
		odd=false
	}
}
func evenprint(){
	if even{
		fmt.Println(i++)
		
	}
}
func main(){
	wg := sync.WaitGroup{}

	wg.Add(2)

	go evenprint()

	go oddPrint()

	wg.Wait()

	fmt.Println("------------------------done---------------------")

}