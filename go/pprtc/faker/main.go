package main

import (
	"fmt"
	"platform-svc-ofac_04_jan/utils/secure"
	"time"
	"github.com/jaswdr/faker"
	"encoding/json"
	"io/ioutil"
)

type Model struct {
	ID        string    `db:"id" json:"id"`
	Entity    string    `db:"entity" json:"entity"`
	Scope     Scope     `db:"scope" json:"scope"`
	Attribute string    `db:"attribute" json:"attribute"`
	Value     string    `db:"value" json:"value"`
	MatchType MatchType `db:"match_type" json:"match_type"`
	Threshold float64   `db:"threshold" json:"threshold"`
	Action    Action    `db:"action" json:"action"`
}

type Scope int

const (
	DEFAULT_SCOPE Scope = iota
	GLOBAL
	REGIONAL
	BIN_SPONSOR
	PROGRAM
	WALLET
	CARD
	USER
)

// Action enum
type Action int

const (
	REJECT_CREATE Action = iota
	ACCEPTED
	CLEARED
)

// MatchType enum
type MatchType int

const (
	STRICT MatchType = iota
	FUZZY
)

type key_Param struct{
	Attribute string    `db:"attribute" json:"attribute"`
	Value     string    `db:"value" json:"value"`
}

m := map[string][]key_Param{
	"DEFAULT_SCOPE":[]{{"Dname","Dval"},{},{}},
	"GLOBAL":[]{{},{},},
	"REGIONAL":[]{{},{},{},},
	"BIN_SPONSOR":[]{{}},
	"PROGRAM":[]{},
	"WALLET":[]{},
	"CARD":[]{},
	"USER":[]{},
}

func GenerateRandomHash(hashLength int) string {
	unixTime := fmt.Sprintf("%d", time.Now().Unix())
	return fmt.Sprintf(
		"%s%s",
		unixTime,
		secure.RandomString(hashLength-len(unixTime), []rune(secure.RuneAlNumCS)),
	)
}
indexm := map[string][string]string{"index": {"_index": "bank_accounts"}}
 
func main(){
	fake := faker.New()

	entity := [...]string{
		"customer",
		"Bank accounts",
		"Beneficiaries",
		"corporations",
		"countries",
		"Addresses",
	}

	for i := 0; i < 100; i++ {
		Model{
			ID : GenerateRandomHash(32),
			Entity: fake.RandomStringElement(entity), 
			Scope: fake.IntBetween(0,8),
			keyp,_ = m[Entity]
			val = fake.RandomStringElement(keyp)
			Attribute : val.Attribute,
			Value:  val.Value,
			MatchType :fake.IntBetween(0,1),
			Threshold:fake.Float64(1, 0, 1),
			Action:fake.IntBetween(0,3),
			
		}
	}
	
		file, _ := json.MarshalIndent(Model, "", " ")
 
		_ = ioutil.WriteFile("test.json", file, 0777)//UGO permission  RWE
}
