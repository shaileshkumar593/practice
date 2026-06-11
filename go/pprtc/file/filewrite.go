package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line."}
	for _, line := range linesToWrite {
		file.WriteString(line + "\n")
	}
}

/*
It does not maintain any buffer and writes to the file immediately as soon as
the write is called.*/
