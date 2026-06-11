package main


import "fmt"

func Sum(nums... int)int{
	total := 0

	for _, n := range nums{
		total += n
	}

	return total
}

func main(){
	fmt.Println(Sum(1,2))
	fmt.Println(Sum(5,8,9,7))

	arr := []int{3,7,8,9}

	fmt.Println(Sum(arr...))
}