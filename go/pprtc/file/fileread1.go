/*
We will use ioutil to read the file directly. The ioutil package makes sure
that the file gets closed after reading. So, we don’t need to worry about
the memory leakage due to not closing the file.
*/

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("file1.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	fmt.Println(string(content)) // This is some content
}
