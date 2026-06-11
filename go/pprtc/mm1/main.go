package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://127.0.0.1:8083/api/v1/businesses?hash_id=c4960c56240e4bda976f1bf700451459&provider_account_number=Test6789&is_active=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-MM-Request-ID", "001")
	req.Header.Add("X-MM-Sub-Request-ID", "002")
	req.Header.Add("X-MM-Tenant-Code", "phatomph")
	req.Header.Add("X-MM-Auth-Token", "NzI4OTE5MmNiOGQ5Mjg5MWUwY2NlZTY2OTcwOThmNDU6OGRxMk5RVzNQNk1GbUtlWXlNRUtqZGpoUEhMalJLWXM0TjdhRWh1Tg==")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
