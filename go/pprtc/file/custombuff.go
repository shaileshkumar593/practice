/*
We can also create a custom buffer size for our writer using bufio.NewWriterSize().
In the below example, we create a write with a custom buffer size of 10 bytes.
If you carefully notice the bytes “Available” and “Buffered” in output, you will
find that writes to file keep happening in between when the buffer is full and at
the end when writer.Flush() is called.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Write file using bufio writer
func main() {
	file, err := os.Create("./temp  .txt")
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriterSize(file, 10)
	linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line."}
	for _, line := range linesToWrite {
		bytesWritten, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
		}
		fmt.Printf("Bytes Written: %d\n", bytesWritten)
		fmt.Printf("Available: %d\n", writer.Available())
		fmt.Printf("Buffered : %d\n", writer.Buffered())
	}
	writer.Flush()
}
