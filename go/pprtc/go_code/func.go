package main
import "fmt"

func area(length,width int)int{
ar := length*width 
return ar
}

func main(){
fmt.Printf("Area of Rectangle : %d",area(12,20))
}

