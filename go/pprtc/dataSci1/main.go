package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"os"
)

func main() {
	// Load data from CSV file
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	df := dataframe.ReadCSV(file)

	// Display summary statistics of the data
	fmt.Println(df.Describe())
}
