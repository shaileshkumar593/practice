package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://beta-api.mmvpay.com/sgmcrpo/v1/oauth/consumer/WgfzcWvUm5ChOS25OYzSG0oYiMWlkgdd/prefund/credit"
	method := "POST"

	payload := strings.NewReader("amount=1&transaction_ref_id=test678787999999999999")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-MM-Auth-Token", "V2dmemNXdlVtNUNoT1MyNU9ZelNHMG9ZaU1XbGtnZGQ6YW5aeWowaWhDaDhNMDRvU3NJQ3RBeDBnUUJmZzh3bUt2WWx2MFM2MWlCRmllUTEzWTcyTFhpOGthRHVHUlpyeA==")
	req.Header.Add("Authorization", "Basic V2dmemNXdlVtNUNoT1MyNU9ZelNHMG9ZaU1XbGtnZGQ6YW5aeWowaWhDaDhNMDRvU3NJQ3RBeDBnUUJmZzh3bUt2WWx2MFM2MWlCRmllUTEzWTcyTFhpOGthRHVHUlpyeA==")
	req.Header.Add("User-Agent", "Apidog/1.0.0 (https://www.apidog.com)")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println(req)
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
