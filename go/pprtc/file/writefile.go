package main

import (
	"io/ioutil"
	"log"
)

func main() {
	linesToWrite := "This is an example to show how to write to file using ioutil"
	err := ioutil.WriteFile("temp.txt", []byte(linesToWrite), 0777)
	if err != nil {
		log.Fatal(err)
	}
}

/*
ioutil.WriteFile() is kind of a shortcut to writing to a file. It takes in three-parameter
– (1) file (2) The string to be written (3) Permission mode of the file. The third parameter is
used to create a file with that permission if the file doesn’t already exist. One calling
ioutil.WriteFile(), it will Create the file if not present with the specified permission
Write to the file Close the file See below example: If temp.txt doesn’t exist then
it will create a new temp.txt file with permission 0777*/
