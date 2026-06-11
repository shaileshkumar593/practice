package main

func main() {
	s := []int{8, 7, 6}

	p := []int{44, 66, 99}

	s = append(s, p...)

}
