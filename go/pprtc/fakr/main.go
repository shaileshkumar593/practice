package main

import (
	"fmt"

	"github.com/jaswdr/faker"
)

func main() {

	f := faker.New()
	type ExampleStruct struct {
		SimpleStringField  string
		SimpleNumber       int
		SimpleBool         bool
		SomeFormatedString string    `fake:"??? ###"`
		SomeStringArray    [5]string `fake:"????"`
	}

	example := ExampleStruct{}
	f.Struct().Fill(&example)
	fmt.Printf("%+v", example)

	fmt.Println(f.RandomFloat(1, 80, 90))

	type CreditRequest struct {
		RefNo         string  `json:"refNo" validate:"required"`
		CreditType    string  `json:"creditType" validate:"required"`
		Amount        float64 `json:"amount" validate:"required"`
		Status        string  `json:"status" validate:"required"`
		OriginAccount struct {
			BankName      string `json:"bankName" validate:"required"`
			AccountNumber string `json:"acctNo" validate:"required"`
			AccountName   string `json:"acctName" validate:"required"`
		} `json:"originAcct"`

		DestinationAccount struct {
			AccountNumber string `json:"acctNo" validate:"required"`
			AccountName   string `json:"acctName" validate:"required"`
		} `json:"destinationAcct"`

		TimeStamp string `json:"timestamp" validate:"required"`
	}

	example1 := CreditRequest{}
	f.Struct().Fill(&example1)
	fmt.Printf("%+v", example1)
}
