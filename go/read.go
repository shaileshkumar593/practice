package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter size: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Enter values:")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(arr)
}

/*

https://github.com/krishnaik06/The-Grand-Complete-Data-Science-Materials
https://github.com/cnkyrpsgl/leetcode
https://github.com/milanm/DevOps-Roadmap
https://github.com/devops-by-examples/complete-devops-course
https://github.com/in28minutes/devops-master-class/tree/master/kubernetes
*/
