func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	test()

	fmt.Println("Main continues")
}

func test() {
	panic("boom")
}