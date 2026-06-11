package main 


import "fmt"


func CombinationSum(data []int, target int)[][]int{
	var result [][]int

	var dfs func(start int, remain int, path []int)

	dfs = func(start int, remain int, path []int){
		if remain == 0{
			temp := append([]int{}, path...)
			result = append(result, temp)
			return
		}

		if remain < 0{
			return 
		}

		for i := start; i < len(data);i++{
			path = append(path, data[i])
			dfs(i, remain - data[i], path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, target, []int{})

	return result
}

func main() {
	fmt.Println(CombinationSum([]int{1,2, 3,4,5, 6}, 6))
}