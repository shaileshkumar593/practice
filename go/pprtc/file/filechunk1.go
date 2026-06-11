/*
A more efficient approach when reading a file is to read it by chunks.
That means we don’t take it all completely in memory but instead we load
it by chunking. It’s very important to read large files in chunk to avoid
out of memory errors.
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("file1.txt")
	if err != nil {
		fmt.Println("Error opening file!!!")
	}
	defer file.Close()

	// declare chunk size
	const maxSz = 4

	// create buffer
	b := make([]byte, maxSz)

	for {
		// read content to buffer
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(string(b[:readTotal])) // print content from buffer
	}
}
