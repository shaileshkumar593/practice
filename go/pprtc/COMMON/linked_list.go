package main

import (
	"fmt"
	"os"
)

type Node struct {
	data int   //
	link *Node //
}

var first1 Node
var cnt int = 0

func inserfront(ele int) {
	var temp Node

	temp.data = ele
	temp.link = nil
	fmt.Println("node   ", temp)

	if &first1 == nil {
		fmt.Println("List is empty  ")
		first1 = temp
		cnt = cnt + 1
	} else {
		temp.link = &first1
		first1 = temp
		cnt = cnt + 1
	}
}

func display() {
	if cnt == 0 {
		fmt.Println("List is empty ")
	} else {
		cur := &first1

		for cur != nil {
			fmt.Print(cur.data, "--->")
			cur = cur.link
		}
	}

}

func deleteNode() {
	var node_cnt int
	fmt.Println("Enter node count to delete ")
	fmt.Scanf("%d", &node_cnt)

	if cnt < node_cnt {
		fmt.Println("Node count not available")
	} else {
		cur := &first1
		cnt1 := 1 //
		for cur != nil {
			if cnt1 == node_cnt {
				break
			}
			cnt1++
			cur = cur.link
		}
	}

}

func main() {
	var ch int
	var ele int
	// var ele_slice = make([]int, 5, 10)
	// for i := 0; i < len(ele_slice); i++ {
	// 	_, err := fmt.Scanf("%d", &ele_slice[i])
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	// for _, val := range ele_slice {
	// 	fmt.Println(val)
	// }

	for {
		fmt.Println("Enter  choice for 1. inserfront 2. display 3. deleteNode 4.exit")
		fmt.Scanf("%d", &ch)
		if ch == 1 {
			fmt.Print("Enter Element to insert ")
			fmt.Scanf("%d", &ele)
		}
		//fmt.Printf("%T", ch)
		//fmt.Println()
		switch ch {
		case 1:

			inserfront(ele)

		case 2:
			display()

		case 3:
			deleteNode()

		case 4:
			os.Exit(0)

		}
	}

}
