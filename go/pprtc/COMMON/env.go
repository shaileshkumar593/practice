package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("FOO", "1,2,3,4")
	v := os.Getenv("FOO")

	fmt.Println("-----------", v)
	fmt.Println()
	/* for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	} */
}
