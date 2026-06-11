// You can edit this code!
// Click here and start typing.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	indexm := map[string]map[string]string{"index": {"_index": "bank_accounts"}}
	data := map[string]int{
		"a": 1,
		"b": 2,
	}

	i, err := json.MarshalIndent(indexm, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	fmt.Println(string(i))

	file, err := os.OpenFile("example.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Could not open example.txt")
		return
	}

	defer file.Close()

	_ = ioutil.WriteFile("test.json", i, 0777)
	_ = ioutil.WriteFile("test.json", b, 0777)
}
