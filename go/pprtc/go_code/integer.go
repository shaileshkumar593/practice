package main
import "fmt"

func main(){
// 8 bit unsigned integer 
var x uint8 = 225
fmt.Println(x+1,x)

// 16 bit signed number
var y int16 = 32767
fmt.Println(y+5,y)

var t rune = 3455555
fmt.Println(t+1,t)
}