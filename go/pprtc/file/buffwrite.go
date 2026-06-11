/*
bufio package can be used to do a buffered write in Go. It has a default buffer
size of 4096 bytes however a custom buffer size can also be specified.
The main advantage of buffered write is that it keeps
the data to be written in the buffer and thus minimizes the number of times we
have to do the IO operation. It does a write when writer.flush() is called
Buffer is full on next write. For example, let’s say it has 3500 bytes in full in
buffer and buffer has a size of 4096. The next line to be written is of 1000 bytes.
Since it 3500 + 1000 > 4096, it will write some of the bytes to file so that some
buffer is available for next write.If there is no buffer then there will be IO
write on every line. Let’s see an example of a write to a file. We are using
writer.Available() and writer.Buffered() to print the available and used buffer
size respectively.
Write happens at the end when writer.Flush() is called as buffer never really
gets full in between
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("./temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
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
