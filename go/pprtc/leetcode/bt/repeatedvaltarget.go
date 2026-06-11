package main 


import (
	"fmt"
	"sort"
)

func combinationSum(arry []int, targetSum int)[][]int{
	sort.Ints(arry)

	var result [][]int


	var dfs  func(start int, remain int, path []int)

	dfs = func(start int, remain int, path []int){
		if remain == 0{
			temp := append([]int{}, path...)
			result = append(result,temp)
			return 
		}


		for i := start; i < len(arry); i++{
			if i > start && arry[i] == arry[i-1]{
				continue
			}

			if arry[i] > remain{
				break
			}

			path = append(path, arry[i])

			dfs(i+1, remain - arry[i], path)
			path = path[:len(path) - 1]
		}
	}

	dfs(0, targetSum, []int{})

	return result
}

func main(){
	fmt.Println(combinationSum2(
		[]int{10, 1, 2, 7, 6, 1, 5},
		8,
	))

}

