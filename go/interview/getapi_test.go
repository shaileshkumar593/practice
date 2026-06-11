package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func Test_GetUrldata(t *testing.T) {

	ctx := context.Background()
	url := "https://gorest.co.in/public/v2/users"
	resp, _ := GetUrlData(ctx, url)

	if resp.StatusCode == http.StatusOK {
		fmt.Println("get request success")
	}

}
