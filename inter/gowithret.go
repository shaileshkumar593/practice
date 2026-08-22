pacakge main 


import (
	"fmt"
)

func Message() string{
	fmt.Println("hello")

	return "World"
}

func main(){
	go Message()// error 
}