package main

import (
	"fmt"
)

type User struct {
	Id        int
	Name      string
	OrderList []int
}

type OrderStruct struct {
	Id    int
	Name  string
	Price float32
}

type PricePaidByCust struct {
	TotalPrice float32
	Name       string
}

func getDetailsById1(id int, val []any) any {

	for _, v := range val {

		t := reflect.TypeOf(v)

		if t == reflect.TypeOf(User{}) {
			u := v.(User)
			if u.Id == id {
				return u
			}
		}

		if t == reflect.TypeOf(OrderStruct{}) {
			o := v.(OrderStruct)
			if o.Id == id {
				return o
			}
		}
	}

	return nil
}

func getDetailsById(id int, val []any) any {

	for _, v := range val {

		switch vs := v.(type) {

		case User:
			if vs.Id == id {
				return vs
			}

		case OrderStruct:
			if vs.Id == id {
				return vs
			}
		}
	}

	return nil
}

func getTotalPricePaidByEachCust(c []User, u []OrderStruct) []PricePaidByCust {

	pricelist := make([]PricePaidByCust, 0)
	for _, val := range c {
		total := float32(0.0)
		for _, id := range val.OrderList {
			for _, uval := range u {
				if uval.Id == id {
					total += uval.Price
				}
			}

		}
		pl := PricePaidByCust{
			TotalPrice: total,
			Name:       val.Name,
		}
		pricelist = append(pricelist, pl)

	}

	return pricelist
}
func main() {
	fmt.Println("Hello, World!")

	inventory := []OrderStruct{{
		Id:    24,
		Name:  "orange",
		Price: 40.50,
	}, {
		Id:    44,
		Name:  "Apple",
		Price: 70.20,
	}, {
		Id:    83,
		Name:  "Grapes",
		Price: 70.50,
	},
	}

	Customer := []User{
		{
			Id:        247235,
			Name:      "Sohan",
			OrderList: []int{44, 83},
		}, {
			Id:        247835,
			Name:      "Mohan",
			OrderList: []int{44, 83, 24},
		}, {
			Id:        287235,
			Name:      "Rohan",
			OrderList: []int{24, 83},
		},
	}

		custdetail := getDetailsById(247235, Customer).(User)

		inventorydetail := getDetailsById(44, inventory).(OrderStruct)

		fmt.println("customer ", custdetail)
		fmt.Println("Inventory ", inventorydetail)

	pl := getTotalPricePaidByEachCust(Customer, inventory)
	fmt.Println(pl)

}
