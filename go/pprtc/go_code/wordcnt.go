/*Steps in solving problem

1. Read filename passed from  command line
2. Readfile file
3. Convert into Lowercase
4. Match Regular Expression with file
5. Store String and their count


*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

//declaring Structure for storing key and count for each
type keyval struct {
	key string
	val int
}

func main() {

	filetoread := os.Args[1]
	NoofWord := 20
	//Regular Expression to read anykind of letter from any Language
	//MustCompile parses the RegularExpression with Panic mode

	reg := regexp.MustCompile(`\p{Ll}+`)

	//Read the file and log the fatal error
	bs, err := ioutil.ReadFile(filetoread)

	if err != nil {
		log.Fatal(err)
	}
	//convert file text to lowercase and compare all succesive matches find in the file
	// Return the slice of matches found.
	text := strings.ToLower(string(bs))
	matches := reg.FindAllString(text, -1)

	// store the matched string with count in  map
	groups := make(map[string]int)
	for _, match := range matches {
		groups[match]++

	}

	/*
		When iterating over a map with a range loop, the iteration order is not specified and is not guaranteed to be the same from one iteration to the next.
		If you require a stable iteration order you must maintain a separate data structure that specifies that order*/

	var keyvals []keyval
	for k, v := range groups {
		keyvals = append(keyvals, keyval{k, v})
	}

	// Sort the value based on count
	sort.Slice(keyvals, func(i, j int) bool {
		return keyvals[i].val > keyvals[j].val
	})
	fmt.Println("Index  Word  Frequency")
	fmt.Println("====  ====  =========")
	for index := 1; index <= NoofWord; index++ {
		word := keyvals[index-1].key
		freq := keyvals[index-1].val
		fmt.Printf("%2d    %-4s    %5d\n", index, word, freq)
	}

}
