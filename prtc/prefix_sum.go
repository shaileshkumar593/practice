package main 

import(
	"fmt"
)

func prefixSum(arr []int)[]int{
	prefix := make([]int, len(arr))

	prefix[0] = arr[0]

	for i := 1; i < len(arr); i++{
		prefix[i] = prefix[i-1] + arr[i]
	}

	return prefix
}

func main(){
	arr := []int{2,8,9,7,3,66}
	fmt.Println(prefixSum(arr))
}