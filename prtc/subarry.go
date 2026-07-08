package main 


import(
	"fmt"
)

func subarraySum(nums []int, target int) int{
	count := 0
	sum := 0

	mp := make(map[int][]int)

	mp[0] =[]int{-1}

	for i, num := range nums{
		sum += num

		if indices, ok := mp[sum -target]; ok {
			for _, start := range indices{
				fmt.Println(nums[start+1:i+1])
				count++
			}
		}

		mp[sum] = append(mp[sum],i)
	}

    return count
}

func main() {

	arr := []int{1, 1, 1}

	count := subarraySum(arr, 2)

	fmt.Println("Count:", count)
}
