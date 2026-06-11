package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Model struct {
	ID          int64  `db:"id"`
	HashID      string `db:"hash_id" json:"hash_id,omitempty"`
	ProgramCode string `db:"program_code" json:"program_code,omitempty"`
	Scope       int    `db:"scope"`
	ConfigKey   string `db:"config_Key" json:"config_Key,omitempty"`
	ConfigValue string `db:"config_value" json:"config_value,omitempty"`
	Priority    int8   `db:"priority"`
}

func main() {
	q := Model{
		ID:          2,
		HashID:      "errr",
		ProgramCode: "pppp",
		Scope:       3,
		ConfigKey:   "ertwu",
		ConfigValue: "powert",
		Priority:    3,
	}
	fmt.Println(reflect.TypeOf(q))
	fmt.Println(reflect.ValueOf(q))
	fmt.Println(reflect.TypeOf(q).Kind())

	m := map[string]interface{}{
		"ID":          2,
		"HashID":      "utrgy",
		"ProgramCode": "tryu",
		"Scope":       4,
		"ConfigKey":   "rety",
		"ConfigValue": "yuptr",
		"Priority":    8,
	}
	// Convert map to json string
	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	// Convert json string to struct
	var me Model
	if err := json.Unmarshal(jsonStr, &me); err != nil {
		fmt.Println(err)
	}
	fmt.Println("1", me.HashID)
	fmt.Println("2", me.ConfigKey)
	fmt.Println("3", me.ConfigValue)
	fmt.Println("4", me.ID)
	fmt.Println("5", me.Priority)
	fmt.Println("6", me.ProgramCode)
	fmt.Println("7", me.Scope)
}
