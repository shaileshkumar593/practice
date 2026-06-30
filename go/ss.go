package main


import(
	"fmt"
)

func main(){
	var a [5]int =[5]int{3,6,2,9,5}
	var b = [...]int{9,2,7,8}
	s := a[1:3]
	fmt.Println(s,len(s),cap(s))
	fmt.Println(b,len(b),cap(b))

	fmt.Println(a)
}