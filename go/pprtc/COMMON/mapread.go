package main

import (
	"fmt"
)

var Config_Level = map[string]int{
	"GLOBAL":       1,
	"NATION":       2,
	"REGIONAL":     3,
	"BIN_SPONSOR ": 4,
	"PROGRAM":      5,
	"WALLET":       6,
	"CARD":         7,
	"USER":         8,
}

func main() {
	scope := "NATION"

	fmt.Println(Config_Level[scope])
}
